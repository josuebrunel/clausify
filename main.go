package clausify

import (
	"errors"
	"strconv"
	"strings"
)

// ErrInvalidOperator describes an invalid operator error
var ErrInvalidOperator = errors.New("Invalid operator")

// Concat concatenate strings
func concat(ss ...string) string {
	var sb strings.Builder
	for _, s := range ss {
		sb.WriteString(s)
	}
	return sb.String()
}

// Clausifier is an interface that wraps a basic Clausify method
type Clausifier interface {
	Clausify(k string, vv []string) (Condition, error)
}

// Condition describes a SQL Clause condition
type Condition struct {
	Expression string
	Variables  []interface{}
}

// Clause describe a SQL Where clause
type Clause struct {
	Conditions string
	Variables  []interface{}
}

// AddCondition adds a where clause condition to the current where clause
func (c *Clause) AddCondition(cond Condition) {
	if c.Conditions == "" {
		c.Conditions = cond.Expression
	} else {
		c.Conditions = concat(c.Conditions, " AND ", cond.Expression)
	}
	c.Variables = append(c.Variables, cond.Variables...)
}

var operators = map[string]string{
	"eq": "=", "neq": "!=",
	"gt": ">", "gte": ">=",
	"lt": "<", "lte": "<=",
	"in": "IN", "nin": "NOT IN",
	"like": "LIKE", "ilike": "ILIKE", "nlike": "NOT LIKE",
	"between": "BETWEEN", "nbetween": "NOT BETWEEN",
}

// QSClausifier is the default clausifier
type QSClausifier struct {
	Separator   string
	Placeholder string
	Operators   map[string]string
}

// GetOperator returns the operator key
func (c QSClausifier) GetOperator(k string) (string, string) {
	op := strings.Split(k, c.Separator)
	if len(op) == 2 {
		return op[0], op[1]
	}
	return k, "eq"
}

// BuildCondition return condition variables with the right type
func (c QSClausifier) BuildCondition(k string, o string, v string) Condition {
	cond := Condition{}
	var nv []interface{}
	for _, e := range strings.Split(v, ",") {
		if val, err := strconv.Atoi(e); err == nil {
			nv = append(nv, val)
			continue
		}
		nv = append(nv, e)
	}
	if strings.Contains(o, "IN") {
		cond.Expression = concat(k, " ", o, " ", c.Placeholder)
		cond.Variables = append(cond.Variables, nv)
	} else {
		if strings.Contains(o, "BETWEEN") {
			cond.Expression = concat(k, " ", o, " ", c.Placeholder, " AND ", c.Placeholder)
		} else {
			cond.Expression = concat(k, " ", o, " ", c.Placeholder)
		}
		cond.Variables = append(cond.Variables, nv...)
	}
	return cond
}

// Clausify is the
func (c QSClausifier) Clausify(k string, vv []string) (Condition, error) {
	cond := Condition{}
	k, op := c.GetOperator(k)
	if _, in := c.Operators[op]; !in {
		return cond, ErrInvalidOperator
	}
	return c.BuildCondition(k, c.Operators[op], vv[0]), nil
}

// With tuns url.Query into where clause condtion by passing a custom operator
func With(q map[string][]string, cf Clausifier) (Clause, error) {
	c := Clause{}
	for k, v := range q {
		cond, err := cf.Clausify(k, v)
		if err != nil {
			return c, err
		}
		c.AddCondition(cond)
	}
	return c, nil
}

// Clausify takes an url.Query and turns it into a Where clause conditions
func Clausify(q map[string][]string) (Clause, error) {
	return With(q, QSClausifier{
		Separator: "__", Operators: operators, Placeholder: "?"})
}

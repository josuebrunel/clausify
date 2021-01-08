package clausify

import (
	"errors"
	"strconv"
	"strings"
)

const separator string = "__"

type opfunc func(k string, vv []string) (c Condition)

// ErrInvalidOperator describes an invalid operator error
var ErrInvalidOperator = errors.New("Invalid operator")

// OPExpression describe expression of an operator
type OPExpression struct {
	Name          string
	Expression    string
	NumExpression string
}

// Condition describes a where clause condition
type Condition struct {
	Expression string
	Variables  []interface{}
}

func isNumeric(v string) bool {
	if _, err := strconv.Atoi(v); err != nil {
		return false
	}
	return true
}

func concat(ss ...string) string {
	var sb strings.Builder
	for _, s := range ss {
		sb.WriteString(s)
	}
	return sb.String()
}

func op(o OPExpression, k string, vv []string) (c Condition) {
	c = Condition{}
	if strings.Contains(o.Name, "between") {
		vv = strings.Split(vv[0], ",")
	}
	for _, v := range vv {
		if isNumeric(v) {
			c.Expression = concat(k, o.NumExpression)
		} else {
			c.Expression = concat(k, o.Expression)
		}
		if strings.Contains(o.Name, "in") {
			if len(c.Variables) == 0 {
				c.Variables = append(c.Variables, vv)
			}
		} else {
			c.Variables = append(c.Variables, v)
		}
	}
	return c
}

func eq(k string, vv []string) (c Condition) {
	return op(OPExpression{Expression: " = '?'", NumExpression: " = ?"}, k, vv)
}

func neq(k string, vv []string) (c Condition) {
	return op(OPExpression{Expression: " != '?'", NumExpression: " != ?"}, k, vv)
}

func gt(k string, vv []string) (c Condition) {
	return op(OPExpression{NumExpression: " > ?"}, k, vv)
}

func gte(k string, vv []string) (c Condition) {
	return op(OPExpression{NumExpression: " >= ?"}, k, vv)
}

func lt(k string, vv []string) (c Condition) {
	return op(OPExpression{NumExpression: " < ?"}, k, vv)
}

func lte(k string, vv []string) (c Condition) {
	return op(OPExpression{NumExpression: " <= ?"}, k, vv)
}

func like(k string, vv []string) (c Condition) {
	return op(OPExpression{Expression: " LIKE '?'"}, k, vv)
}

func ilike(k string, vv []string) (c Condition) {
	return op(OPExpression{Expression: " ILIKE '?'"}, k, vv)
}

func nlike(k string, vv []string) (c Condition) {
	return op(OPExpression{Expression: " NOT LIKE '?'"}, k, vv)
}

func in(k string, vv []string) (c Condition) {
	return op(OPExpression{Name: "in", Expression: " IN (?)"}, k, vv)
}

func nin(k string, vv []string) (c Condition) {
	return op(OPExpression{Name: "nin", Expression: " NOT IN (?)"}, k, vv)
}

func between(k string, vv []string) (c Condition) {
	return op(OPExpression{Name: "between", Expression: " BETWEEN '?' AND '?'", NumExpression: " BETWEEN ? AND ?"}, k, vv)
}

func nbetween(k string, vv []string) (c Condition) {
	return op(OPExpression{Name: "nbetween", Expression: " NOT BETWEEN '?' AND '?'", NumExpression: " NOT BETWEEN ? AND ?"}, k, vv)
}

func getOperator(key string) (string, string) {
	op := strings.Split(key, separator)
	if len(op) == 2 {
		return op[0], op[1]
	}
	return key, "eq"
}

var operators = map[string]opfunc{
	"eq":       eq,
	"neq":      neq,
	"gt":       gt,
	"gte":      gte,
	"lt":       lt,
	"lte":      lte,
	"like":     like,
	"ilike":    ilike,
	"nlike":    nlike,
	"in":       in,
	"nin":      nin,
	"between":  between,
	"nbetween": nbetween,
}

// Clause describe a WHERE Clause and its conditions
type Clause struct {
	Conditions string
	Variables  []interface{}
}

// AddCondition add a clause condition
func (c *Clause) AddCondition(cond Condition) {
	if c.Conditions == "" {
		c.Conditions = cond.Expression
	} else {
		c.Conditions = concat(c.Conditions, " AND ", cond.Expression)
	}
	for v := range cond.Variables {
		c.Variables = append(c.Variables, v)
	}
}

// Clausify takes an url.Query and turns it into a Where clause conditions
func Clausify(q map[string][]string) (Clause, error) {
	c := Clause{}
	for k, v := range q {
		k, op := getOperator(k)
		if _, ok := operators[op]; !ok {
			return c, ErrInvalidOperator
		}
		c.AddCondition(operators[op](k, v))
	}
	return c, nil
}

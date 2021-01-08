package operators

import (
	"github.com/josuebrunel/clausify/clause"
	"strconv"
	"strings"
)

// Operator is an interface describing a where clause operator
type Operator interface {
	Lookup(key string, values []string) (clause.Condition, error)
}

const separator string = "__"

type opfunc func(k string, vv []string) (c clause.Condition)

func isNumeric(v string) bool {
	if _, err := strconv.Atoi(v); err != nil {
		return false
	}
	return true
}

// OPExpression describe expression of an operator
type OPExpression struct {
	Name          string
	Expression    string
	NumExpression string
}

// ClausifyOperator is the operator implemeting clausify operations
type ClausifyOperator struct {
	Separator  string
	Operations map[string]opfunc
}

// Lookup is a func extracting the operation and returning the clause condition
func (co ClausifyOperator) Lookup(k string, vv []string) (c clause.Condition, err error) {
	err = nil
	k, op := getOperator(k)
	if _, ok := operators[op]; !ok {
		return c, clause.ErrInvalidOperator
	}
	return co.Operations[op](k, vv), nil
}

func op(o OPExpression, k string, vv []string) (c clause.Condition) {
	c = clause.Condition{}
	if strings.Contains(o.Name, "between") {
		vv = strings.Split(vv[0], ",")
	}
	for _, v := range vv {
		if isNumeric(v) {
			c.Expression = clause.Concat(k, o.NumExpression)
		} else {
			c.Expression = clause.Concat(k, o.Expression)
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

func eq(k string, vv []string) (c clause.Condition) {
	return op(OPExpression{Expression: " = '?'", NumExpression: " = ?"}, k, vv)
}

func neq(k string, vv []string) (c clause.Condition) {
	return op(OPExpression{Expression: " != '?'", NumExpression: " != ?"}, k, vv)
}

func gt(k string, vv []string) (c clause.Condition) {
	return op(OPExpression{NumExpression: " > ?"}, k, vv)
}

func gte(k string, vv []string) (c clause.Condition) {
	return op(OPExpression{NumExpression: " >= ?"}, k, vv)
}

func lt(k string, vv []string) (c clause.Condition) {
	return op(OPExpression{NumExpression: " < ?"}, k, vv)
}

func lte(k string, vv []string) (c clause.Condition) {
	return op(OPExpression{NumExpression: " <= ?"}, k, vv)
}

func like(k string, vv []string) (c clause.Condition) {
	return op(OPExpression{Expression: " LIKE '?'"}, k, vv)
}

func ilike(k string, vv []string) (c clause.Condition) {
	return op(OPExpression{Expression: " ILIKE '?'"}, k, vv)
}

func nlike(k string, vv []string) (c clause.Condition) {
	return op(OPExpression{Expression: " NOT LIKE '?'"}, k, vv)
}

func in(k string, vv []string) (c clause.Condition) {
	return op(OPExpression{Name: "in", Expression: " IN (?)"}, k, vv)
}

func nin(k string, vv []string) (c clause.Condition) {
	return op(OPExpression{Name: "nin", Expression: " NOT IN (?)"}, k, vv)
}

func between(k string, vv []string) (c clause.Condition) {
	return op(OPExpression{Name: "between", Expression: " BETWEEN '?' AND '?'", NumExpression: " BETWEEN ? AND ?"}, k, vv)
}

func nbetween(k string, vv []string) (c clause.Condition) {
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

// DefaultOperator is an instance of ClausifyOperator
var DefaultOperator = ClausifyOperator{
	Separator:  "__",
	Operations: operators,
}

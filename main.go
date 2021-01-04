package clausify

import (
	"errors"
	"strconv"
	"strings"
)

const separator string = "__"

type opfunc func(k, v string) string

// ErrInvalidOperator describes an invalid operator error
var ErrInvalidOperator = errors.New("Invalid operator")

// Operator describes operator experessions
type Operator struct {
	Expression    string
	NumExpression string
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

func op(o Operator, k, v string) (s string) {
	if isNumeric(v) {
		s = concat(k, o.NumExpression)
	} else {
		s = concat(k, o.Expression)
	}
	return
}

func eq(k, v string) (s string) {
	return op(Operator{Expression: " = '?'", NumExpression: " = ?"}, k, v)
}

func neq(k, v string) (s string) {
	return op(Operator{Expression: " != '?'", NumExpression: " != ?"}, k, v)
}

func gt(k, v string) (s string) {
	return op(Operator{NumExpression: " > ?"}, k, v)
}

func gte(k, v string) (s string) {
	return op(Operator{NumExpression: " >= ?"}, k, v)
}

func getOperator(key string) (string, string) {
	op := strings.Split(key, separator)
	if len(op) == 2 {
		return op[0], op[1]
	}
	return key, "eq"
}

var operators = map[string]opfunc{
	"eq":  eq,
	"neq": neq,
	"gt":  gt,
	"gte": gte,
}

// Clause describe a WHERE Clause statement
type Clause struct {
	Statement string
	Variables []interface{}
}

// AddCondition add a clause condition
func (c *Clause) AddCondition(s string, v interface{}) {
	if c.Statement == "" {
		c.Statement = s
	} else {
		c.Statement = concat(c.Statement, " AND ", s)
	}
}

// Clausify takes an url.Query and turns it into an SQL Statement
func Clausify(q map[string][]string) (Clause, error) {
	c := Clause{}
	for k, v := range q {
		k, op := getOperator(k)
		if _, ok := operators[op]; !ok {
			return c, ErrInvalidOperator
		}
		c.AddCondition(operators[op](k, v[0]), v[0])
		c.Variables = append(c.Variables, v[0])
	}
	return c, nil
}

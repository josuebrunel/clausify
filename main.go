package clausify

import (
	"strconv"
	"strings"
)

const seperator string = "__"

type opfunc func(k, v string) string

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

func eq(k, v string) (s string) {
	if isNumeric(v) {
		s = concat(k, " = ?")
	} else {
		s = concat(k, " = '?'")
	}
	return
}

func neq(k, v string) (s string) {
	if isNumeric(v) {
		s = concat(k, " != ?")
	} else {
		s = concat(k, " != '?'")
	}
	return
}

func getOperator(key string) (string, string) {
	op := strings.Split(key, seperator)
	if len(op) == 2 {
		return op[0], op[1]
	}
	return key, "eq"
}

var operators = map[string]opfunc{
	"eq":  eq,
	"neq": neq,
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
func Clausify(q map[string][]string) Clause {
	c := Clause{}
	for k, v := range q {
		k, op := getOperator(k)
		c.AddCondition(operators[op](k, v[0]), v[0])
		c.Variables = append(c.Variables, v[0])
	}
	return c
}

package clause

import (
	"errors"
	"strings"
)

// ErrInvalidOperator describes an invalid operator error
var ErrInvalidOperator = errors.New("Invalid operator")

// Concat concatenate strings
func Concat(ss ...string) string {
	var sb strings.Builder
	for _, s := range ss {
		sb.WriteString(s)
	}
	return sb.String()
}

// Condition describes a where clause condition
type Condition struct {
	Expression string
	Variables  []interface{}
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
		c.Conditions = Concat(c.Conditions, " AND ", cond.Expression)
	}
	for v := range cond.Variables {
		c.Variables = append(c.Variables, v)
	}
}

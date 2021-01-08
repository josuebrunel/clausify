package clausify

import (
	"github.com/josuebrunel/clausify/clause"
	"github.com/josuebrunel/clausify/operators"
)

// With tuns url.Query into where clause condtion by passing a custom operator
func With(q map[string][]string, op operators.Operator) (clause.Clause, error) {
	c := clause.Clause{}
	for k, v := range q {
		cond, err := op.Lookup(k, v)
		if err != nil {
			return c, err
		}
		c.AddCondition(cond)
	}
	return c, nil
}

// Clausify takes an url.Query and turns it into a Where clause conditions
func Clausify(q map[string][]string) (clause.Clause, error) {
	return With(q, operators.DefaultOperator)
}

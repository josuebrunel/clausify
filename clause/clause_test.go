package clause

import (
	"github.com/matryer/is"
	"strings"
	"testing"
)

func TestConcat(t *testing.T) {
	is := is.New(t)
	is.Equal(Concat("Hello", " world", " !"), "Hello world !")
}

func TestClause(t *testing.T) {
	is := is.New(t)
	c := Clause{}
	cond1 := Condition{Expression: "username = '?'", Variables: []interface{}{"loking"}}
	cond2 := Condition{Expression: "id > ?", Variables: []interface{}{1000}}
	c.AddCondition(cond1)
	c.AddCondition(cond2)
	is.True(strings.Contains(c.Conditions, "username = '?'"))
	is.True(strings.Contains(c.Conditions, "id > ?"))
	is.Equal(len(c.Variables), 2)
}

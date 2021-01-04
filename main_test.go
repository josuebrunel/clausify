package clausify

import (
	"github.com/matryer/is"
	"strings"
	"testing"
)

type Values map[string][]string

func TestInvalidOperator(t *testing.T) {
	is := is.New(t)
	v := Values{"username__xy": []string{"josh"}}
	_, err := Clausify(v)
	is.Equal(err.Error(), "Invalid operator")
}

func TestEqual(t *testing.T) {
	is := is.New(t)
	v := Values{"username": []string{"josh"}}
	c, _ := Clausify(v)
	is.Equal(c.Statement, "username = '?'")
	is.Equal(len(c.Variables), 1)
	v = Values{"username": []string{"josh"}, "age": []string{"30"}}
	c, _ = Clausify(v)
	is.Equal(strings.Contains(c.Statement, "username = '?'"), true)
	is.Equal(strings.Contains(c.Statement, "age = ?"), true)
	is.Equal(len(c.Variables), 2)
}

func TestNotEqual(t *testing.T) {
	is := is.New(t)
	v := Values{"username__neq": []string{"josh"}}
	c, _ := Clausify(v)
	is.Equal(c.Statement, "username != '?'")
	is.Equal(len(c.Variables), 1)
	v = Values{"username__neq": []string{"josh"}, "age__neq": []string{"30"}}
	c, _ = Clausify(v)
	is.Equal(strings.Contains(c.Statement, "username != '?'"), true)
	is.Equal(strings.Contains(c.Statement, "age != ?"), true)
	is.Equal(len(c.Variables), 2)
}

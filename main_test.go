package clausify

import (
	"github.com/matryer/is"
	"strings"
	"testing"
)

type Values map[string][]string

func TestEqual(t *testing.T) {
	is := is.New(t)
	v := Values{"username": []string{"josh"}}
	c := Clausify(v)
	is.Equal(c.Statement, "username = '?'")
	is.Equal(len(c.Variables), 1)
	v = Values{"username": []string{"josh"}, "age": []string{"30"}}
	c = Clausify(v)
	is.Equal(strings.Contains(c.Statement, "username = '?'"), true)
	is.Equal(strings.Contains(c.Statement, "age = ?"), true)
	is.Equal(len(c.Variables), 2)
}

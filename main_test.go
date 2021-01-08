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

func TestGreaterThan(t *testing.T) {
	is := is.New(t)
	v := Values{"price__gt": []string{"15"}, "name": []string{"book"}}
	c, _ := Clausify(v)
	is.True(strings.Contains(c.Statement, "price > ?"))
	is.True(strings.Contains(c.Statement, "name = '?'"))
	is.Equal(len(c.Variables), 2)
}

func TestGreaterThanEqual(t *testing.T) {
	is := is.New(t)
	v := Values{"price__gte": []string{"15"}, "name": []string{"book"}}
	c, _ := Clausify(v)
	is.True(strings.Contains(c.Statement, "price >= ?"))
	is.True(strings.Contains(c.Statement, "name = '?'"))
	is.Equal(len(c.Variables), 2)
}

func TestLessThan(t *testing.T) {
	is := is.New(t)
	v := Values{"price__lt": []string{"15"}, "name": []string{"book"}}
	c, _ := Clausify(v)
	is.True(strings.Contains(c.Statement, "price < ?"))
	is.True(strings.Contains(c.Statement, "name = '?'"))
	is.Equal(len(c.Variables), 2)
}

func TestLessThanEqual(t *testing.T) {
	is := is.New(t)
	v := Values{"price__lte": []string{"15"}, "name": []string{"book"}}
	c, _ := Clausify(v)
	is.True(strings.Contains(c.Statement, "price <= ?"))
	is.True(strings.Contains(c.Statement, "name = '?'"))
	is.Equal(len(c.Variables), 2)
}

func TestLike(t *testing.T) {
	is := is.New(t)
	v := Values{"price__lte": []string{"15"}, "name__like": []string{"book"}}
	c, _ := Clausify(v)
	is.True(strings.Contains(c.Statement, "price <= ?"))
	is.True(strings.Contains(c.Statement, "name LIKE '?'"))
	is.Equal(len(c.Variables), 2)
}

func TestILike(t *testing.T) {
	is := is.New(t)
	v := Values{
		"price__lte":  []string{"15"},
		"name__ilike": []string{"book"},
		"category":    []string{"fruits"},
	}
	c, _ := Clausify(v)
	is.True(strings.Contains(c.Statement, "price <= ?"))
	is.True(strings.Contains(c.Statement, "name ILIKE '?'"))
	is.Equal(len(c.Variables), 3)
}

func TestNotLike(t *testing.T) {
	is := is.New(t)
	v := Values{"price__lte": []string{"15"}, "name__nlike": []string{"book"}}
	c, _ := Clausify(v)
	is.True(strings.Contains(c.Statement, "price <= ?"))
	is.True(strings.Contains(c.Statement, "name NOT LIKE '?'"))
	is.Equal(len(c.Variables), 2)
}

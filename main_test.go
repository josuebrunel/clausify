package clausify

import (
	"github.com/josuebrunel/clausify/clause"
	"github.com/matryer/is"
	"net/url"
	"strings"
	"testing"
)

func getURLQuery(uri string) map[string][]string {
	u, _ := url.Parse(uri)
	return u.Query()
}

func TestInvalidOperator(t *testing.T) {
	is := is.New(t)
	q := getURLQuery("https://httpbin.org/?username__xy=josh")
	_, err := Clausify(q)
	is.Equal(err.Error(), "Invalid operator")
}

func TestEqual(t *testing.T) {
	is := is.New(t)
	q := getURLQuery("https://httpbin.org/?username=josh")
	c, _ := Clausify(q)
	is.Equal(c.Conditions, "username = '?'")
	is.Equal(len(c.Variables), 1)
	q = getURLQuery("https://httpbin.org/?username=josh&age=30")
	c, _ = Clausify(q)
	is.Equal(strings.Contains(c.Conditions, "username = '?'"), true)
	is.Equal(strings.Contains(c.Conditions, "age = ?"), true)
	is.Equal(len(c.Variables), 2)
}

func TestNotEqual(t *testing.T) {
	is := is.New(t)
	q := getURLQuery("https://httpbin.org/?username__neq=josh")
	c, _ := Clausify(q)
	is.Equal(c.Conditions, "username != '?'")
	is.Equal(len(c.Variables), 1)
	q = getURLQuery("https://httpbin.org/?username__neq=josh&age__neq=30")
	c, _ = Clausify(q)
	is.Equal(strings.Contains(c.Conditions, "username != '?'"), true)
	is.Equal(strings.Contains(c.Conditions, "age != ?"), true)
	is.Equal(len(c.Variables), 2)
}

func TestGreaterThan(t *testing.T) {
	is := is.New(t)
	q := getURLQuery("https://httpbin.org/?price__gt=15&name=book")
	c, _ := Clausify(q)
	is.True(strings.Contains(c.Conditions, "price > ?"))
	is.True(strings.Contains(c.Conditions, "name = '?'"))
	is.Equal(len(c.Variables), 2)
}

func TestGreaterThanEqual(t *testing.T) {
	is := is.New(t)
	q := getURLQuery("https://httpbin.org/?price__gte=15&name=book")
	c, _ := Clausify(q)
	is.True(strings.Contains(c.Conditions, "price >= ?"))
	is.True(strings.Contains(c.Conditions, "name = '?'"))
	is.Equal(len(c.Variables), 2)
}

func TestLessThan(t *testing.T) {
	is := is.New(t)
	q := getURLQuery("https://httpbin.org/?price__lt=15&name=book")
	c, _ := Clausify(q)
	is.True(strings.Contains(c.Conditions, "price < ?"))
	is.True(strings.Contains(c.Conditions, "name = '?'"))
	is.Equal(len(c.Variables), 2)
}

func TestLessThanEqual(t *testing.T) {
	is := is.New(t)
	q := getURLQuery("https://httpbin.org/?price__lte=15&name=book")
	c, _ := Clausify(q)
	is.True(strings.Contains(c.Conditions, "price <= ?"))
	is.True(strings.Contains(c.Conditions, "name = '?'"))
	is.Equal(len(c.Variables), 2)
}

func TestLike(t *testing.T) {
	is := is.New(t)
	q := getURLQuery("https://httpbin.org/?price__lte=15&name__like=book")
	c, _ := Clausify(q)
	is.True(strings.Contains(c.Conditions, "price <= ?"))
	is.True(strings.Contains(c.Conditions, "name LIKE '?'"))
	is.Equal(len(c.Variables), 2)
}

func TestILike(t *testing.T) {
	is := is.New(t)
	q := getURLQuery("https://httpbin.org/?price__lte=15&name__ilike=book&category=fruits")
	c, _ := Clausify(q)
	is.True(strings.Contains(c.Conditions, "price <= ?"))
	is.True(strings.Contains(c.Conditions, "name ILIKE '?'"))
	is.Equal(len(c.Variables), 3)
}

func TestNotLike(t *testing.T) {
	is := is.New(t)
	q := getURLQuery("https://httpbin.org/?price__lte=15&name__nlike=book")
	c, _ := Clausify(q)
	is.True(strings.Contains(c.Conditions, "price <= ?"))
	is.True(strings.Contains(c.Conditions, "name NOT LIKE '?'"))
	is.Equal(len(c.Variables), 2)
}

func TestIn(t *testing.T) {
	is := is.New(t)
	q := getURLQuery("https://httpbin.org/?id__in=2,4,6")
	c, _ := Clausify(q)
	is.True(strings.Contains(c.Conditions, "id IN (?)"))
	is.Equal(len(c.Variables), 1)
}

func TestNotIn(t *testing.T) {
	is := is.New(t)
	q := getURLQuery("https://httpbin.org/?id__nin=2,4,6")
	c, _ := Clausify(q)
	is.True(strings.Contains(c.Conditions, "id NOT IN (?)"))
	is.Equal(len(c.Variables), 1)
}

func TestBetween(t *testing.T) {
	is := is.New(t)
	q := getURLQuery("https://httpbin.org/?category=fruits&price__between=10,20")
	c, _ := Clausify(q)
	is.True(strings.Contains(c.Conditions, " BETWEEN ? AND ?"))
	is.Equal(len(c.Variables), 3)
}

func TestNotBetween(t *testing.T) {
	is := is.New(t)
	q := getURLQuery("https://httpbin.org/?category=fruits&price__nbetween=10,20")
	c, _ := Clausify(q)
	is.True(strings.Contains(c.Conditions, " NOT BETWEEN ? AND ?"))
	is.Equal(len(c.Variables), 3)
}

type MyOperator struct {
	Separator string
}

func (m MyOperator) Lookup(k string, vv []string) (clause.Condition, error) {
	op := strings.Split(k, m.Separator)
	if op[1] == "<>" {
		return clause.Condition{
			Expression: clause.Concat(op[0], " <> ?"),
			Variables:  []interface{}{vv[0]},
		}, nil
	}
	return clause.Condition{}, nil
}

func TestCustomOperator(t *testing.T) {
	is := is.New(t)
	q := getURLQuery("https://httpbin.org/?id-<>=1")
	c, _ := With(q, MyOperator{Separator: "-"})
	is.True(strings.Contains(c.Conditions, "id <> ?"))
	is.Equal(len(c.Variables), 1)
}

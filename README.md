[![test](https://github.com/josuebrunel/clausify/workflows/test/badge.svg)](https://github.com/josuebrunel/clausify/actions?query=workflow%3Atest)
[![coverage](https://coveralls.io/repos/github/josuebrunel/clausify/badge.svg?branch=main)](https://coveralls.io/github/josuebrunel/clausify?branch=main)
[![goreportcard](https://goreportcard.com/badge/github.com/josuebrunel/clausify)](https://goreportcard.com/report/github.com/josuebrunel/clausify)
[![gopkg](https://pkg.go.dev/badge/github.com/josuebrunel/clausify.svg)](https://pkg.go.dev/github.com/josuebrunel/clausify)
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/josuebrunel/clausify/blob/master/LICENSE)

# Clausify

*Clausify* helps you turn you *url query strings* into *SQL Where clause*
It supports SQL Comparison operators and some logical operators

## Installation

```go
go get github.com/josuebrunel/clausify
```

## Quickstart

```go

import (
    "github.com/josuebrunel/clausify"
    "net/url"
    "fmt"
)

u, _ := url.Parse("https://httpbin.org/?email__like=@toto.com&age__gte=24&company=toto")

c, err := clausify.Clausify(u.Query())
if err != nil {
    // do whatever
}

fmt.Printf("%s\n", c.Conditions) // email like '?' AND age >= ? AND company = '?'
fmt.Printf("%v\n", c.Variables) // ["@toto.com", 24, "toto"]
```

## Supported operators by the operators.DefaultOperator

| Query string filters                      | SQL Operators                                   |
|-------------------------------------------|-------------------------------------------------|
| element=value                             | element **=** value OR element = 'value'        |
| element__neq=value                        | element **!=** value OR element != 'value'      |
| element__gt=value                         | element **>** value                             |
| element__gte=value                        | element **>=** value                            |
| element__lt=value                         | element **<** value                             |
| element__lte=value                        | element **<=** value                            |
| element__like=value                       | element **LIKE** 'value'                        |
| element__ilike=value                      | element **ILIKE** 'value'                       |
| element__nlike=value                      | element **NOT LIKE** 'value'                    |
| element__in=value1,value2,valueN          | element **IN** (value1, value2, valueN)         |
| element__nin=value1,value2,valueN         | element **NOT IN** (value1, value2, valueN)     |
| element__between=left,right               | element **BETWEEN** left **AND** right          |
| element__nbetween=left,right              | element **NOT BETWEEN** left **AND** right      |

## Implement a custom Operator

Use a struct implementing Clausifier interface with **Clausify** method as below

```go
import (
    "github.com/josuebrunel/clausify"
    "net/url"
    "strings"
    "testing"
    "errors"
)

type MyClausifier struct {
	Separator string
}

func (m MyClausifier) Clausify(k string, vv []string) (clause.Condition, error) {
	op := strings.Split(k, m.Separator)
	if op[1] == "<>" {
		return clause.Condition{
			Expression: clause.Concat(op[0], " <> ?"),
			Variables:  []interface{}{vv[0]},
		}, nil
	}
	return clause.Condition{}, errors.New("Invalid operator")
}

u, _ := url.Parse("https://httpbin.org/?id-<>=1")
q := u.Query()
c, _ := clausify.With(q, MyClausifier{Separator: "-"})
is.True(strings.Contains(c.Conditions, "id <> ?"))
is.Equal(len(c.Variables), 1)
```

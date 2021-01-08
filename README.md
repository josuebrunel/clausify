[![test](https://github.com/josuebrunel/clausify/workflows/test/badge.svg)](https://github.com/josuebrunel/clausify/actions?query=workflow%3Atest)
[![coverage](https://coveralls.io/repos/github/josuebrunel/clausify/badge.svg?branch=main)](https://coveralls.io/github/josuebrunel/clausify?branch=main)
[![goreportcard](https://goreportcard.com/badge/github.com/josuebrunel/clausify)](https://goreportcard.com/report/github.com/josuebrunel/clausify)
[![gopkg](https://pkg.go.dev/badge/github.com/josuebrunel/clausify.svg)](https://pkg.go.dev/github.com/josuebrunel/clausify)
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/josuebrunel/clausify/blob/master/LICENSE)

# Clausify

*Clausify* helps you turn you *url query strings* into *SQL Where clause statement*
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

fmt.Printf("%s\n", c.Statement) // email like '?' AND age >= ? AND company = '?'
fmt.Printf("%v\n", c.Variables) // ["@toto.com", 24, "toto"]
```

## Supported operators

| Query string filters                      | SQL Operator                                    |
|-------------------------------------------|-------------------------------------------------|
| <element>=<value>                         | <element> **=** <value> OR <element> = '<value>'    |
| <element>__neq=<value>                      | <element> **!=** <value> OR <element> != '<value>'  |
| <element>__gt=<value>                      | <element> > <value>                             |
| <element>__gte=<value>                    | <element> >= <value>                            |
| <element>__lt=<value>                     | <element> < <value>                             |
| <element>__lte=<value>                    | <element> <= <value>                            |
| <element>__like=<value>                   | <element> LIKE '<value>'                        |
| <element>__ilike=<value>                  | <element> ILIKE '<value>'                       |
| <element>__nlike=<value>                  | <element> NOT LIKE '<value>'                    |
| <element>__in=<value1>,<value2>,<valueN>  | <element> IN (<value1>, <value2>, <valueN>)     |
| <element>__nin=<value1>,<value2>,<valueN> | <element> NOT IN (<value1>, <value2>, <valueN>) |
| <element>__between=<left>,<right>         | <element> BETWEEN <left> AND <right>            |
| <element>__nbetween=<left>,<right>        | <element> NOT BETWEEN <left> AND <right>        |


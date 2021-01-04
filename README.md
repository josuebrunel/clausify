[![test](https://github.com/josuebrunel/clausify/actions?query=workflow%3Atest)](https://github.com/josuebrunel/clausify/workflows/test/badge.svg)
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
    "fmt"
)

qs := r.Url.Values // ?email__like=@toto.com&age__gte=24&company=toto

c, _ := clausify.Clausify(qs)

fmt.Printf("%s\n", c.Statement) // email like '?' AND age >= ? AND company = '?'
fmt.Printf("%v\n", c.Variables) // ["@toto.com", 24, "toto"]
```

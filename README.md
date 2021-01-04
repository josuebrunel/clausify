![test](https://github.com/josuebrunel/clausify/workflows/test/badge.svg)

# Clausify

*Clausify* helps you turn you *url query strings* into *SQL Where clause statement*

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

c := clausify.Clausify(qs)

fmt.Printf("%s\n", c.Statement) // email like '?' AND age >= ? AND company = '?'
fmt.Printf("%v\n", c.Variables) // ["@toto.com", 24, "toto"]
```

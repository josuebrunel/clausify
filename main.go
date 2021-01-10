package main

import (
	"clause/clause"
	"fmt"
)

func main() {
	url := "https://httpbin.org/journeys/?toto=hello&departure_datetime__gte=2021Z&departure_place=Paris&free__gt=1&arrival_datetime=2021-01-09T"
	r := clause.Clausify(url)
	fmt.Println(r)
}

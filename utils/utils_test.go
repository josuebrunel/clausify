package utils_test

import (
	"clause/utils"
	"testing"
)

func TestTokenize(t *testing.T) {
	s := "https://httpbin.org/journeys/?toto=hello&departure_datetime__gte=2021Z&departure_place=Paris&free__gt=1&arrival_datetime=2021-01-09T"
	v := "?"
	e := []string{"https://httpbin.org/journeys/", "toto=hello&departure_datetime__gte=2021Z&departure_place=Paris&free__gt=1&arrival_datetime=2021-01-09T"}
	r := utils.Tokenize(s, v)
	i := 0

	if len(e) != len(r) {
		t.Fatal("Error with the tokenizer, go : ", r)
	}

	for i < len(r) {
		if e[i] == r[i] {
			i++
			continue
		}
		t.Fatal("Error with the tokenizer, got : ", r)
	}
}

func TestReplacePatternsInString(t *testing.T) {
	s := "jacques"
	p := []string{"ac", "ques"}
	r := "osh"
	w := "joshosh"
	v := utils.ReplacePatternsInString(s, p, r)

	if w != v {
		t.Fatal("Error with the Pattern replacer, got : ", v)
	}
}

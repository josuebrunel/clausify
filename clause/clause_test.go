package clause_test

import (
	"clause/clause"
	"testing"
)

type test struct {
	input    string
	expected string
}

func TestClausify(t *testing.T) {
	tests := []test{
		{input: "https://httpbin.org/journeys/?toto=hello&departure_datetime__gte=2021Z&departure_place=Paris&free__gt=1&arrival_datetime=2021-01-09T", expected: "SELECT * FROM ___ WHERE toto = hello AND departure_datetime >= 2021Z AND departure_place = Paris AND free > 1 AND arrival_datetime = 2021-01-09T"},
		{input: "https://httpbin.org/?email__like=@toto.com&age__gte=24&company=toto", expected: "SELECT * FROM ___ WHERE email like @toto.com AND age >= 24 AND company = toto"},
		{input: "https://httpbin.org/?email__like=@toto.com&age__gte=24&?list_a=1&list_a=2&list_a=3&list_b=1&list_b=2&list_b=3&list_c=1,2,3", expected: "SELECT * FROM ___ WHERE email like @toto.com AND age >= 24 AND list_a = 1 AND list_a = 2 AND list_a = 3 AND list_b = 1 AND list_b = 2 AND list_b = 3 AND list_c = 1,2,3"},
		{input: "title=Query_string&action=edit", expected: "SELECT * FROM ___ WHERE title = Query_string AND action = edit"},
		{input: "https://example.com/path/to/page?name=ferret&color=purple", expected: "SELECT * FROM ___ WHERE name = ferret AND color = purple"},
		{input: "https://httpbin.org/journeys/?toto=hello&departure_datetime__gte=2021-01-08T23:42:34+01:00&departure_place=Paris&free__gt=1&arrival_datetime=2021-01-09T", expected: "SELECT * FROM ___ WHERE toto = hello AND departure_datetime >= 2021-01-08T23:42:34+01:00 AND departure_place = Paris AND free > 1 AND arrival_datetime = 2021-01-09T"},
		{input: "http://www.localhost.com/Webform2.aspx?name=Atilla&lastName=Ozgur", expected: "SELECT * FROM ___ WHERE name = Atilla AND lastName = Ozgur"},
		{input: "http://www.localhost.com/Webform2.aspx?cost__between=10,100", expected: "SELECT * FROM ___ WHERE cost between 10 AND 100"},
		{input: "http://www.localhost.com/Webform2.aspx?destination=Paris&cost__between=10,100", expected: "SELECT * FROM ___ WHERE destination = Paris AND cost between 10 AND 100"},
		{input: "http://www.localhost.com/v1?destination=Paris&cost__between=10,100&email__like=@toto.com&age__gte=24", expected: "SELECT * FROM ___ WHERE destination = Paris AND cost between 10 AND 100 AND email like @toto.com AND age >= 24"},
	}

	for _, tc := range tests {
		actual := clause.Clausify(tc.input)

		if actual != tc.expected {
			t.Fatal("Error with the Clausifier, got : ", actual)
		}
	}
}

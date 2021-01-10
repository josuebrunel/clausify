package clause

import (
	"clause/operators"
	"clause/utils"
	"strings"
)

const (
	queryParamSep     string = "__"
	fieldnameValueSep string = "="
	pairSeriesSep     string = "&"
	placeholderSQL    string = "SELECT * FROM ___ WHERE "
	placeholderAND    string = " AND "
	emptyString       string = " "
	queryParamStart   string = "?"
)

// GenerateSQLQuery generates the actual SQL query
func GenerateSQLQuery(g map[int][]string) string {
	var s string

	/* dfs through the Tree
	   notice a list of visited nodes isn't necessary here
	   as the traversal is like a linear pattern
	   the produced tree is always linear (no recursion)
	*/
	for i := 0; i < len(g); i++ {
		v := g[i]

		if i == 0 {
			/* edge case for between and not between */
			if v[1] == "between" || v[1] == "nbetween" {
				s = placeholderSQL
				var u string
				u = v[0] + operators.GetOperator(v[1]) + v[2]
				s += strings.ReplaceAll(u, ",", placeholderAND)
				continue
			}
			s += placeholderSQL +
				v[0] +
				operators.GetOperator(v[1]) +
				v[2]
			continue
		}

		/* edge case for between and not between */
		if v[1] == "between" || v[1] == "nbetween" {
			var u string
			u = v[0] + operators.GetOperator(v[1]) + v[2]
			s += placeholderAND + strings.ReplaceAll(u, ",", placeholderAND)
			continue
		}
		/* edge case for */
		s += placeholderAND + v[0] + operators.GetOperator(v[1]) + v[2]
	}

	return s
}

// Process function processes the querystrings
func Process(s string) []string {
	var t []string
	var r string
	var slc = []string{queryParamSep,
		fieldnameValueSep,
	}

	if !strings.Contains(s, queryParamStart) {
		if !strings.Contains(s, queryParamSep) {
			r = strings.ReplaceAll(s, fieldnameValueSep, " = ")
		} else {
			r = utils.ReplacePatternsInString(s, slc, emptyString)
		}
	} else {
		r = strings.ReplaceAll(s, queryParamStart, emptyString)
		u := utils.Tokenize(r, emptyString)
		r = u[1]
		if !strings.Contains(s, queryParamSep) {
			r = strings.ReplaceAll(r, fieldnameValueSep, " = ")
		} else {
			r = utils.ReplacePatternsInString(r, slc, emptyString)
		}
	}

	t = utils.Tokenize(r, emptyString)

	return t
}

// Clausify function is the entry point
func Clausify(qs string) string {
	ret := utils.Tokenize(qs, pairSeriesSep)

	/* trie structure */
	g := map[int][]string{}
	i := 0
	/* leaves number are all 3 to fullfil exp1 ops exp2 constraint */
	r := make([]string, 3)

	/* generate a Trie */
	for _, v := range ret {
		r = Process(v)
		g[i] = r
		i++
	}

	q := GenerateSQLQuery(g)

	return q
}

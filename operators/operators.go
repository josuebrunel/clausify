package operators

var operators = map[string]string{
	"=":        " = ",
	"eq":       " = ",
	"neq":      " != ",
	"gt":       " > ",
	"gte":      " >= ",
	"lt":       " < ",
	"lte":      " <= ",
	"like":     " like ",
	"ilike":    " ilike ",
	"in":       " in ",
	"nin":      "not in",
	"nlike":    " not like ",
	"between":  " between ",
	"nbetween": " not between ",
}

// GetOperator returns the operator as a string
func GetOperator(k string) string {
	return operators[k]
}

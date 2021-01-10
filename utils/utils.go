package utils

import "strings"

// Tokenize tokenizes the string
func Tokenize(s string, t string) []string {
	return strings.Split(s, t)
}

// ReplacePatternsInString replace multiple patterns in a string
func ReplacePatternsInString(s string, p []string, r string) string {
	for _, v := range p {
		s = strings.ReplaceAll(s, v, r)
	}
	return s
}

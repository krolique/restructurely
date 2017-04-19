package main

import (
	"strings"
)

// tokenizes a given string (value) into string tokens created (split) by
// the delimiter parameter passed in to the function
func tokenize(value string, delimiter rune) []string {
	// this function defines how each token is determined when used by the
	// FieldsFunc
	splitter := func(r rune) bool {
		switch r {
		case delimiter:
			return true
		}
		return false
	}

	// The output of FieldsFunc could be an array with one element or as
	// many elements as there are tokens.
	return strings.FieldsFunc(value, splitter)
}

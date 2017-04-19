package main

import (
	"strings"
)

// tokenizes a directory name into a
func tokenize(value string, delimiter rune) []string {
	// this function defines how each token is determined when used by the
	// FieldsFunc
	splitter := func(r rune) bool {
		switch r {
		// potential to refine the "case" to include multiple case as
		// case '<', '>', ':'
		case delimiter:
			return true
		}
		return false
	}

	return strings.FieldsFunc(value, splitter)
}

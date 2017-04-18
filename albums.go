package main

import (
	"bytes"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const year_pattern = "(19|20)\\d{2}"

var title_replacer = strings.NewReplacer("(", "[", ")", "]", "-", "")

// Returns year found in the string
func find_year(file_name string) (year int, err error) {

	r, err := regexp.Compile(year_pattern)
	if err != nil {
		return -1, err
	}

	matches := r.FindAllString(file_name, -1)

	if err != nil {
		return -1, err
	}

	if len(matches) == 0 {
		return -1, errors.New("Not a single year pattern found.")
	}

	/* If we've encountered multiple year matches this is something that
	 * requires user input (for the time being) and should result in
	 * termination of execution. */
	if len(matches) != 1 {
		return -1, errors.New("Multiple year patterns found.")
	}

	converted, err := strconv.Atoi(matches[0])
	if err != nil {
		return -1, err
	}

	return converted, nil
}

func album_name_sanitizer(file_name string) (sanitized string) {

	file_name = title_replacer.Replace(file_name)

	var buffer bytes.Buffer
	for _, token := range strings.Fields(file_name) {
		switch token {
		case "2CD":
		case "[]":
		case "[FLAC]":
		default:
			buffer.WriteString(token + " ")
		}
	}

	return strings.Title(strings.TrimSpace(buffer.String()))
}

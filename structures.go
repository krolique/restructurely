package main

import (
	"fmt"
	"strconv"
	"strings"
)

// An Album grouping
type Album struct {
	// The year the Album was published
	year int

	// The title of the Album (display)
	title string
}

func (a *Album) Display() (display_string string) {
	return fmt.Sprintf("<Album title[%s] year [%d]>", a.title, a.year)
}

func (a *Album) FilePath() (file_path string) {
	return fmt.Sprintf("%s (%d)", a.title, a.year)
}

func (a *Album) FromString(value string) error {

	matched_year, err := find_year(value)
	if err != nil {
		return err
	}
	a.year = matched_year
	// if we've found a year then this token can be removed from the string
	value = strings.Replace(value, strconv.Itoa(a.year), "", 1)
	a.title = album_name_sanitizer(value)

	return nil
}

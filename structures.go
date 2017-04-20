package main

import (
	"fmt"
	"strconv"
	"strings"
)

// This data type represents a directory item before it becomes
// an album
type AlbumDirectory struct {
	// root of the album directory
	root string
	// name of the directory without the root
	name string
	// tokens produced from splitting the name of the directory by one of
	// the accepted delimiters
	tokens []string
	// absolute path to the album directory
	full_path string
	// suggested name to rename the directory to
	suggested_name string
}

func (a *AlbumDirectory) Display() string {
	return fmt.Sprintf("<AlbumDirectory name[%s]>", a.name)
}

// An Album grouping
type Album struct {
	// The year the Album was published
	year int
	// The title of the Album (display)
	title string
}

func (a *Album) Display() (display_string string) {
	return fmt.Sprintf("<Album title[%s] year[%d]>", a.title, a.year)
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

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// tokenizes a directory name into a
func tokenize_directory(dir_name string) ([]string, error) {
	// this function defines how each token is determined when used by the
	// FieldsFunc
	splitter := func(r rune) bool {
		switch r {
		// potential to refine the "case" to include multiple case as
		// case '<', '>', ':'
		case '<':
			return true
		}
		return false
	}

	tokens := strings.FieldsFunc(dir_name, splitter)

	if len(tokens) < 1 {
		return nil, errors.New("Failed to tokenize the directory name by any of the known delimiters")
	}

	return tokens, nil
}

// performs pre scan operations on the directory, by laying out the initial
func directory_scan(file_path string) (err error) {

	directories, err := ioutil.ReadDir(file_path)

	if err != nil {
		return err
	}

	for _, directory := range directories {

		tokens, err := tokenize_directory(directory.Name())
		if err == nil {
			for token := range tokens {
				fmt.Println(token)
			}
		}

		full_path := filepath.Join(file_path, directory.Name())
		fmt.Println(full_path)
	}

	return nil
}

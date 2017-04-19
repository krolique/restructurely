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
		case '-':
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

// Scans the folder defined by the `file_path` variable and returns a list
// of possible album directories
func scan(file_path string) ([]AlbumDirectory, error) {

	directories, err := ioutil.ReadDir(file_path)

	if err != nil {
		return nil, err
	}

	var album_dirs []AlbumDirectory

	for _, directory := range directories {

		tokens, err := tokenize_directory(directory.Name())
		if err == nil {
			for token := range tokens {
				fmt.Println(token)
			}
		}

		var new_dir AlbumDirectory
		new_dir.root = file_path
		new_dir.name = directory.Name()
		new_dir.tokens = tokens
		new_dir.full_path = filepath.Join(file_path, directory.Name())
		album_dirs = append(album_dirs, new_dir)
	}

	return album_dirs, nil
}

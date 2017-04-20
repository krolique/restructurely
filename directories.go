package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Scans the folder defined by the `file_path` variable and returns a list
// of possible album directories
func scan(file_path string) ([]AlbumDirectory, error) {

	directories, err := ioutil.ReadDir(file_path)

	if err != nil {
		return nil, err
	}

	all_tokens := map[string]int{}
	var album_dirs []AlbumDirectory
	for _, directory := range directories {
		tokens := tokenize(directory.Name(), '-')

		suggested_name := directory.Name()
		if len(tokens) > 1 {
			for _, token := range tokens {
				_, is_present := all_tokens[token]
				if is_present == true {
					all_tokens[token] += 1
				} else {
					all_tokens[token] = 1
				}
			}
			suggested_name = strings.Join(tokens, "")
		}

		var new_dir AlbumDirectory
		new_dir.root = file_path
		new_dir.name = directory.Name()
		new_dir.tokens = tokens
		new_dir.full_path = filepath.Join(file_path, directory.Name())
		new_dir.suggested_name = suggested_name
		album_dirs = append(album_dirs, new_dir)
	}

	for value, count := range all_tokens {
		if count == len(album_dirs) {
			for i := 0; i < len(album_dirs); i += 1 {
				album_dirs[i].suggested_name = strings.Replace(album_dirs[i].suggested_name, value, "", 1)
			}
		}
	}

	return album_dirs, nil
}

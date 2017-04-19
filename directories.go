package main

import (
	"io/ioutil"
	"path/filepath"
)

// Scans the folder defined by the `file_path` variable and returns a list
// of possible album directories
func scan(file_path string) ([]AlbumDirectory, error) {

	directories, err := ioutil.ReadDir(file_path)

	if err != nil {
		return nil, err
	}

	var album_dirs []AlbumDirectory

	for _, directory := range directories {

		tokens := tokenize(directory.Name(), '-')

		var new_dir AlbumDirectory
		new_dir.root = file_path
		new_dir.name = directory.Name()
		new_dir.tokens = tokens
		new_dir.full_path = filepath.Join(file_path, directory.Name())
		album_dirs = append(album_dirs, new_dir)
	}

	return album_dirs, nil
}

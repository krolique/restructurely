package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func rename(oldpath string, newpath string, peform_rename bool) {

	if peform_rename {
		os.Rename(oldpath, newpath)
	}

	fmt.Println(fmt.Sprintf("Renamed: '%s' to '%s'\n", oldpath, newpath))
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Missing scan directory path as the first command line argument")
		os.Exit(1)
	}

	scan_dir_path := os.Args[1]

	directories, err := scan(scan_dir_path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Restructurely(v0.1.0)\n")
	for _, item := range directories {
		album := new(Album)
		err := album.FromString(item.suggested_name)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Current name: [%s]\nSuggested name: [%s]", item.name, album.FilePath())
		fmt.Print("\nEnter [y] to apply: ")
		reader := bufio.NewReader(os.Stdin)
		response, _, _ := reader.ReadRune()
		if response == 'y' {
			newpath := filepath.Join(item.root, album.FilePath())
			rename(item.full_path, newpath, true)
		}
	}
	fmt.Println("Finished\n")
}

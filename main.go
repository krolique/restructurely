package main

import (
	"fmt"
	"os"
)

func rename(oldpath string, newpath string, peform_rename bool) {

	if peform_rename {
		os.Rename(oldpath, newpath)
	}

	fmt.Println(fmt.Sprintf("Renamed: '%s' to '%s'", oldpath, newpath))
}

func main() {

	scan_dir_path := ""

	if len(os.Args) < 2 {
		fmt.Println("Missing scan directory path as the first command line argument")
		os.Exit(1)
	}

	directories, err := scan(scan_dir_path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for item := range directories {
		fmt.Println(item)
	}
}

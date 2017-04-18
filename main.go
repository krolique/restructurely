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
	path := "D:\\music\\Dishwalla"
	directory_scan(path)
}

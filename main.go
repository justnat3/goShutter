//This is a main.go file. Treat it well. Treat it with kindness. Don't forget to struggle.
package main

import (
	//"encoding/hex"
	"fmt"
	//"log"
	"os"
	"path/filepath"
	//"github.com/devedge/imagehash"
	//"io"
)

var dir string
var file string

//	TODO:1. Be able to scan in a directory and readin each -> image and output -> array hashes
//	     2. take though hashes and either store them in memory or store them in a local db/dump_file
//	     3. File hash for every iterate of the loop

func main() {
	FilePathWalkDir("C:\\git\\test_dir")
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	c := 0
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		c++
		if !info.IsDir() {
			files = append(files, path)
		}
		if c == 1 {
			fmt.Printf("\nDirectory: %s\n", info.Name())
		}
		if c > 1 {
			fmt.Printf("	FileName: %s\n", info.Name())
		}
		return nil
	})
	return files, err
}

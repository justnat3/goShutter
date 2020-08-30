//This is a main.go file. Treat it well. Treat it with kindness. Don't forget to struggle.
package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/devedge/imagehash"
)

var dir string
var file string

//	TODO:1. Be able to scan in a directory and readin each -> image and output -> array hashes
//	     2. take though hashes and either store them in memory or store them in a local db/dump_file
//	     3. File hash for every iterate of the loop

func main() {
	GetFiles("../nameshed")
}

//GrabPNG function using a image hashing library to print out hash values
func GrabPNG(file string) error {

	src, err := imagehash.OpenImg(file)

	if err != nil {
		println(err)
	}

	//Hash x+y and return the hash
	hashLen := 8
	hash, _ := imagehash.Dhash(src, hashLen)

	//print out hash
	fmt.Printf("Your Hash for %s: ", file)
	fmt.Println(hex.EncodeToString(hash))
	return nil
}

//GetFiles : return string -> file for next function
func GetFiles(dir string) string {
	c := 0

	filepath.Walk(dir, func(dir string, info os.FileInfo, err error) error {

		c++
		if c == 1 {
			fmt.Printf("\nDirectory: %s\n", dir)
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		if c > 1 {
			fmt.Printf("INSIDE_FUNC => FileName: %s\n", info.Name())
		}

		file = dir + "/" + info.Name()
		return nil
	})

	if c < 0 {
		println("C == 0i")
	}
	return file
}

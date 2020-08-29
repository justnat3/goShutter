//This is a main.go file. Treat it well. Treat it with kindness. Don't forget to struggle.
package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/devedge/imagehash"
	"github.com/zserge/lorca"
)

var dir string
var file string

//	TODO:1. Be able to scan in a directory and readin each -> image and output -> array hashes
//	     2. take though hashes and either store them in memory or store them in a local db/dump_file
//	     3. File hash for every iterate of the loop

func main() {
	GrabPNG("./")
	ui, err := lorca.New("data:text/html, <h1> Hello world </h1>", "", 480, 320)

	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ui.Done()
}

//GrabPNG function using a image hashing library to print out hash values
func GrabPNG(file string) error {

	src, err := imagehash.OpenImg(file)
	//Hash x+y and return the hash
	hashLen := 8
	hash, _ := imagehash.Dhash(src, hashLen)

	//print out hash
	fmt.Printf("Your Hash for %s: ", file)
	fmt.Println(hex.EncodeToString(hash))
	return nil
}

//GetContentType to grab ext of file for review in the next function
func GetContentType(file string) (string, error) {
	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	// use content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func foo(dir string) string {
	filepath.Walk(dir, func(dir string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err.Error())
		}
		if info.Name() == "." {
			fmt.Println("DIR_START")
		}
		fmt.Printf("FileName: %s\n: ", info.Name())
		file = dir + "/" + info.Name()
		return nil
	})
	return file
}

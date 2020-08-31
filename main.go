//This is a main.go file. Treat it well. Treat it with kindness. Don't forget to struggle.
package main

import (
	"crypto/sha256"
	"encoding/hex"
	//	"encoding/hex"
	"fmt"
	//	"github.com/devedge/imagehash"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type hashed struct {
	path string
	hash string
}

func main() {

	files, err := IOReadDir("C:/Users/Nathan Reed/Desktop/")
	if err != nil {
		panic(err)
	}
	HashFiles(files)
	fmt.Scanln("enter: ")
}

//HashFiles : take array of files and hash them
func HashFiles(files []string) {

	// TODO:
	// Store hashes array in memory.
	// sha256 please

	var readFile string
	var hashes []hashed
	const BlockSize = 64

	for i := 0; i < len(files); i++ {

		readFile = files[i]
		f, err := os.Open(readFile)
		defer f.Close()

		if err != nil {
			log.Fatal("Could not open file")
		}

		buff := make([]byte, 512)
		f.Read(buff)

		switch {
		case http.DetectContentType(buff) == "image/jpeg":

			hasher := sha256.New()
			if _, err := io.Copy(hasher, f); err != nil {
				log.Fatal(err)
			}
			sum := hasher.Sum(nil)
			h := hashed{path: readFile, hash: hex.EncodeToString(sum)}
			hashes = append(hashes, h)
			f.Close()
		case http.DetectContentType(buff) == "image/png":
			hasher := sha256.New()

			if _, err := io.Copy(hasher, f); err != nil {
				log.Fatal(err)
			}
			sum := hasher.Sum(nil)
			h := hashed{path: readFile, hash: hex.EncodeToString(sum)}
			hashes = append(hashes, h)
			f.Close()
		default:
			f.Close()
		}
	}

	for j := 0; j < len(hashes); j++ {
		fmt.Println(hashes[j])
	}
}

//IOReadDir : Read in Directory and spit out file names + PATH
func IOReadDir(root string) ([]string, error) {
	fmt.Println("\n")

	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	c := 0

	if err != nil {
		return files, err
	}
	println("Scanning...  " + root + "/\n")
	for _, file := range fileInfo {
		c++
		filePath := root + file.Name()
		files = append(files, filePath)
	}
	return files, nil
}

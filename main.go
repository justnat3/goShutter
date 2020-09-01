//This is a main.go file. Treat it well. Treat it with kindness. Don't forget to struggle.

//	WHAT IS THIS DOING:
//	Takes a directory -> walks the directory you give it hashing only photo files -> sort and detect duplicates in the path.
//	It then takes the duplicates and sticks them in a directory +1 from your root called dupes/ and you can delete or do whatever at this point.

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type hash struct {
	path string
	hash string
}

func main() {
	files, err := IOReadDir("C:\\Users\\Nathan Reed\\Downloads\\afreightdata\\afreightdata" + "\\")

	if err != nil {
		panic(err)
	}

	HashFiles(files)
	fmt.Scanln("enter: ")
}

//HashFiles : take array of files and hash them
func HashFiles(files []string) {

	var readFile string
	//var hashes []hash
	const BlockSize = 64

	m := make(map[string]string)

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

			_, ok := m[hex.EncodeToString(sum)]
			if ok == false {
				println("False")
			} else {
				println("True")
			}

			m[hex.EncodeToString(sum)] = readFile
			f.Close()
		case http.DetectContentType(buff) == "image/png":
			hasher := sha256.New()

			if _, err := io.Copy(hasher, f); err != nil {
				log.Fatal(err)
			}

			sum := hasher.Sum(nil)

			_, ok := m[hex.EncodeToString(sum)]
			if ok == false {
				println("False")
			} else {
				println("True")
			}
			m[hex.EncodeToString(sum)] = readFile
			f.Close()
		default:
			f.Close()
		}
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

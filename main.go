//This is a main.go file. Treat it well. Treat it with kindness. Don't forget to struggle.
package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/devedge/imagehash"
)

var (
	dir  string
	file string
)

func main() {

	files, err := IOReadDir("C:/Users/Nathan Reed/Desktop")
	if err != nil {
		panic(err)
	}
	HashFiles(files)
	fmt.Scanln("enter: ")
}

//IOReadFile : Take in files from IOREADDir function and read the bytes to check contentType
func HashFiles(files []string) {

	var fileArr []string
	var readFile string

	for i := 0; i < len(files); i++ {

		readFile = files[i]
		buff := make([]byte, 512) // docs tell that it take only first 512 bytes into consideration

		f, err := os.Open(readFile)

		if err != nil {
			log.Fatal("Could not open file")
		}
		f.Read(buff)

		switch {

		case http.DetectContentType(buff) == "image/jpeg":
			src, _ := imagehash.OpenImg(readFile)
			hash, _ := imagehash.Dhash(src, 8)
			//fmt.Println(readFile + ": Success! || Type: " + http.DetectContentType(buff))
			object := readFile + "/" + hex.EncodeToString(hash)
			fileArr = append(fileArr, object)
			f.Close()
		case http.DetectContentType(buff) == "image/png":
			src, _ := imagehash.OpenImg(readFile)
			//fmt.Println(readFile + ": Success! || Type: " + http.DetectContentType(buff))
			hash, _ := imagehash.Dhash(src, 8)
			object := readFile + "/" + hex.EncodeToString(hash)

			fileArr = append(fileArr, object)
			f.Close()
		default:
			f.Close()
		}
	}

	for j := 0; j < len(fileArr); j++ {
		fmt.Println(fileArr[j])
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
		filePath := root + "\\" + file.Name()
		files = append(files, filePath)
	}
	return files, nil
}

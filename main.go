//This is a main.go file. Treat it well. Treat it with kindness. Don't forget to struggle.
package main

import (
	"fmt"
	"os"
	//	"log"
	//"github.com/devedge/imagehash"

	"io/ioutil"
	"net/http"
	"time"
)

var (
	//dir : INITIAL DIRECTORY
	dir string
	//file : files return?
	file string
)

//	TODO:1. Be able to scan in a directory and readin each -> image and output -> array hashes
//	     2. take though hashes and either store them in memory or store them in a local db/dump_file
//	     3. File hash for every iterate of the loop

func main() {

	files, err := IOReadDir("C:/Users/Nathan Reed/Desktop")

	if err != nil {
		println("Did not Preform IOREAD_DIR Function")
	}

	time.Sleep(1 * time.Second)
	IOReadFile(files)
	fmt.Scanln("Enter: ")
}

//HashFiles : Hash Checked Files
func HashFiles(CkFiled string) {

}

//IOReadFile : Take in files from IOREADDir function and read the bytes to check contentType
func IOReadFile(files []string) []string {

	var fileArr []string
	var readFile string

	for i := 0; i < len(files); i++ {

		readFile = files[i]
		buff := make([]byte, 512) // docs tell that it take only first 512 bytes into consideration

		f, err := os.Open(readFile)
		if err != nil {
			fmt.Println("Could not open file")
		}
		f.Read(buff)

		switch {

		case http.DetectContentType(buff) == "image/jpeg":
			fmt.Println(readFile + ": Success! || Type: " + http.DetectContentType(buff))
			fileArr = append(fileArr, readFile)

		case http.DetectContentType(buff) == "image/png":
			fmt.Println(readFile + ": Success! || Type: " + http.DetectContentType(buff))
			fileArr = append(fileArr, readFile)

		default:
			fmt.Println("Failed: " + http.DetectContentType(buff))
		}
	}

	println("outside =>")
	fmt.Println(fileArr)
	println("\n")

	for j := 0; j < len(fileArr); j++ {
		fmt.Println(fileArr[j])
	}
	return nil
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
	println("Scanning...  " + root + "\\\n")
	for _, file := range fileInfo {
		c++
		filePath := root + "\\" + file.Name()
		files = append(files, filePath)
	}
	return files, nil
}

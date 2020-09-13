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

var (
	table = make(map[string]string)
)

func main() {
	fileName, filePath, dupespath, progress, err := IOReadDir("C:\\Users\\Nathan Reed\\Downloads\\afreightdata\\afreightdata\\")
	println(progress)
	if err != nil {
		panic(err)
	}

	HashFiles(fileName, filePath, dupespath)
	fmt.Scanln("enter: ")
}

//HashFiles : take array of files and hash them
func HashFiles(fileName []string, filePath []string, dupespath string) {
	const BlockSize = 64

	for i := 0; i < len(filePath); i++ {
		filePath := filePath[i]
		//fileName := fileName[i]
		//	dupedFile := dupespath + fileName
		f, err := os.Open(filePath)
		defer f.Close()

		if err != nil {
			log.Fatal("Could not open file")
		}

		buff := make([]byte, 512)
		f.Read(buff)

		//somehow creating hashes wrong. perhaps hash, add to map and then iterate through map to find
		// duplicates

		switch {
		case http.DetectContentType(buff) == "image/jpeg":
			hasher := sha256.New()
			if _, err := io.Copy(hasher, f); err != nil {
				log.Fatal(err)
			}
			sum := hasher.Sum(nil)

			key := hex.EncodeToString(sum)
			val, ok := table[key]

			if ok == true {
				println(val)
			} else {
				println("not in table")
			}

			table[key] = filePath

		case http.DetectContentType(buff) == "image/png":
			hasher := sha256.New()
			if _, err := io.Copy(hasher, f); err != nil {
				log.Fatal(err)
			}

			sum := hasher.Sum(nil)

			key := hex.EncodeToString(sum)
			val, ok := table[key]

			if ok == true {
				println(val)
			} else {
				println("not in table")
			}

			table[key] = filePath
			f.Close()
		default:
			f.Close()
		}
	}
}

//IOReadDir : Read in Directory and spit out file names + PATH
func IOReadDir(root string) ([]string, []string, string, int, error) {

	//	var fileObj fileObject
	//	var fileObjects []fileObject
	var fileNames []string
	var filePaths []string

	dupespath := root + "dupes\\"
	if err, _ := os.Stat(dupespath); err == nil {
		os.Mkdir(dupespath, os.FileMode(0522))
	} else {
		log.Println("Already Exists")
	}

	fmt.Println("\n")
	fileInfo, err := ioutil.ReadDir(root)
	c := 0

	if err != nil {
		log.Fatal(err)
	}

	println("Scanning...  " + root + "\\\n")

	for _, file := range fileInfo {
		c++
		fileName := file.Name()
		filePath := root + file.Name()

		fileNames = append(fileNames, fileName)
		filePaths = append(filePaths, filePath)

		//fileObj = fileObject{filePath: root + file.Name(), fileName: file.Name()}
		//fileObjects = append(fileObjects, fileObj)
	}
	progress := len(fileNames)
	return fileNames, filePaths, dupespath, progress, nil
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}

	out, err := os.Create(dst)
	if err != nil {
		in.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	in.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}

	err = out.Sync()
	if err != nil {
		return fmt.Errorf("Sync error: %s", err)
	}

	si, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("Stat error: %s", err)
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return fmt.Errorf("Chmod error: %s", err)
	}

	return nil
}

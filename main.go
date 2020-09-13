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
	"strconv"
	"time"
)

var (
	table = make(map[string]string)
)

func main() {

	fileName, filePath, dupespath, progress, err := IOReadDir("C:\\Users\\Nathan Reed\\Desktop\\")
	itemsInDupes, err := IOReadDupeFolder(dupespath)
	if err != nil {
		panic(err)
	}

	HashFiles(fileName, filePath, dupespath, progress)

	itemsCaught, err := IOReadDupeFolder(dupespath)

	fmt.Printf("Items Caught: %s", strconv.Itoa(itemsCaught-itemsInDupes))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Scanln("enter: ")

}

//HashFiles : take array of files and hash them
func HashFiles(fileName []string, filePath []string, dupespath string, progress int) {

	progressTotal := progress
	start := time.Now()

	for i := 0; i < len(filePath); i++ {
		progress = progress - 1

		fmt.Printf("Progress: [%s/%s]\n", strconv.Itoa(progress), strconv.Itoa(progressTotal))

		filePath := filePath[i]
		fileName := fileName[i]
		dupedFile := dupespath + fileName
		f, err := os.Open(filePath)

		if err != nil {
			log.Fatal("Could not open file")
		}
		defer f.Close()

		buff := make([]byte, 512)
		f.Read(buff)

		switch {

		case http.DetectContentType(buff) == "image/jpeg":

			hasher := sha256.New()
			if _, err := io.Copy(hasher, f); err != nil {
				log.Fatal(err)
			}

			sum := hasher.Sum(nil)
			f.Close()

			key := hex.EncodeToString(sum)
			_, ok := table[key]

			if ok == true {
				err := os.Rename(filePath, dupedFile)
				if err != nil {
					log.Fatal(err)
				}
			}

			table[key] = filePath

		case http.DetectContentType(buff) == "image/png":

			hasher := sha256.New()
			if _, err := io.Copy(hasher, f); err != nil {
				log.Fatal(err)
			}

			f.Close()
			sum := hasher.Sum(nil)

			key := hex.EncodeToString(sum)
			_, ok := table[key]

			if ok == true {
				err := os.Rename(filePath, dupedFile)
				if err != nil {
					log.Fatal(err)
				}
			}

			table[key] = filePath

		default:
			f.Close()
		}
	}

	duration := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n\n", duration)

}

//IOReadDir : Read in Directory and spit out file names + PATH
func IOReadDir(root string) ([]string, []string, string, int, error) {

	var fileNames []string
	var filePaths []string
	var c int = 0

	dupespath := root + "dupes\\"
	if err, _ := os.Stat(dupespath); err == nil {
		os.Mkdir(dupespath, os.FileMode(0522))
	} else {
		log.Println("Already Exists")
	}

	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scanning...  " + root + "\\\n")

	for _, file := range fileInfo {
		c++
		fileName := file.Name()
		filePath := root + file.Name()

		fileNames = append(fileNames, fileName)
		filePaths = append(filePaths, filePath)

	}

	progress := len(fileNames)
	return fileNames, filePaths, dupespath, progress, nil

}

func IOReadDupeFolder(dupespath string) (int, error) {
	var fileNames []string
	var c int = 0

	fileInfo, err := ioutil.ReadDir(dupespath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fileInfo {
		c++
		fileName := file.Name()
		fileNames = append(fileNames, fileName)

	}

	itemsCaught := len(fileNames)
	return itemsCaught, nil

}

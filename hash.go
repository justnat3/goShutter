package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	table = make(map[string]string)
)

//HashFiles : take array of files and hash them
func HashFiles(fileName []string, filePath []string, dupespath string, progress int) {
	// 	1.	Iterate through the files that were read in -> during the ReadDir Phase
	//	2.	Read in the first 512bytes of the file to grab magic bytes
	//	3.	Read into the net/http library to detect the files MimeType
	//	4.	If there is a match we check if there is a match in that location in the table
	//	5. 	if the check returns that it is not in the table we add that key to a location
	//	6.	if the check returns there was a match we put the file in the dupespath

	progressTotal := progress
	start := time.Now()

	for i := 0; i < len(filePath); i++ {
		progress = progress - 1

		fmt.Printf("Progress: [%s/%s]	\n", strconv.Itoa(progress), strconv.Itoa(progressTotal))

		filePath := filePath[i]
		fileName := fileName[i]
		dupedFile := dupespath + fileName
		f, err := os.Open(filePath)

		if err != nil {
			fmt.Println("Error in loop")
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

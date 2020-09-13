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

	"github.com/thatisuday/commando"
)

var (
	dir   string
	table = make(map[string]string)
)

func main() {

	// wd : get working directory for help info section
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	WelcomeStr := "Current Working Directory: " + wd + "\n\nINFO:\n  - This program searches for dupes at the limitations of your drives write speed\n  - If you would like to use this program in your Current working directory. ' goDupe dir . ' is what you are looking for"

	// Commando : Used to get args from user
	commando.
		SetExecutableName("GoDupe").
		SetVersion("v0.5.2").
		SetDescription(WelcomeStr)

	commando.
		Register("dir").
		AddArgument("dir", "Path to desired directory", "").
		SetShortDescription("path to Desired Directory").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			dir = args["dir"].Value
		})

	commando.Parse(nil)

	// if the user-input is "." we can assume that it is the current directory
	// otherwise no dir was specified
	if dir == "." {
		res, err := os.Getwd()
		if err != nil {
			fmt.Printf("\nError: %s", err)
		}
		dir = res
		fmt.Printf("Current Working Directory: %s\n", dir)
	}

	// if user input was not '.' we can assume that what comes next is a directory.
	// if this dir is not valid we can exit and alert the user
	td, err := os.Stat(dir)
	if os.IsNotExist(err) {
		log.Fatal("Folder does not exist.")
		log.Println(td)
	}

	// Read IOReadDir Comment
	fileName, filePath, dupespath, progress, err := IOReadDir(dir)

	//Get Items already in dupes folder. If Dupes does not exist we can create it.
	itemsInDupes, err := IOReadDupeFolder(dupespath)
	if err != nil {
		panic(err)
	}

	// Read Hashing Comment
	HashFiles(fileName, filePath, dupespath, progress)

	//We can assume that the items-caught were not already in the dupes folder so -> substract the total
	itemsCaught, err := IOReadDupeFolder(dupespath)
	fmt.Printf("Items Caught: %s", strconv.Itoa(itemsCaught-itemsInDupes))

	// Do not escape if run from exe
	fmt.Scanln("enter: ")
}

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

//IOReadDir : Read in Directory and spit out file names + PATH
func IOReadDir(root string) ([]string, []string, string, int, error) {

	var fileNames []string
	var filePaths []string
	var c int = 0

	// if dupes path does not exist -> create it
	dupespath := root + "dupes\\"
	if err, _ := os.Stat(dupespath); err == nil {
		os.Mkdir(dupespath, os.FileMode(0522))
	} else {
		log.Println("\n DupesDir: Already Exists")
	}

	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Println("Could not open file\n")
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

//	IOReadDupeFolder : read in what is in the dupesfolder
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

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//IOReadDir : Read in Directory and spit out file names + PATH
func IOReadDir(root string) ([]string, []string, string, int) {
	var fileNames []string
	var filePaths []string
	var dupespath string
	var c int = 0

	// Clean Root Path
	//root = filepath.Clean(root)

	//Decide on OS to permit
	dupespath = root + "dupes\\"
	fmt.Printf("Path to Dupes: %s\n", dupespath)

	// if dupes path does not exist -> create it
	if err, _ := os.Stat(dupespath); err == nil {
		err := os.Mkdir(dupespath, os.FileMode(0522))
		if err != nil {
			fmt.Printf("\nNOTICE:\nConsider adding a \\ or / to the end to complete the path.\n\n")
			log.Fatal(err)
		}
	}

	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		log.Println("Could not open file")
	}

	fmt.Println("Scanning...  " + root + "\n")

	for _, file := range fileInfo {
		c++
		fileName := file.Name()
		filePath := root + file.Name()

		fileNames = append(fileNames, fileName)
		filePaths = append(filePaths, filePath)

	}

	progress := len(fileNames)
	return fileNames, filePaths, dupespath, progress

}

// IOReadDupeFolder : read in what is in the dupesfolder
func IOReadDupeFolder(dupespath string) int {
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
	return itemsCaught
}

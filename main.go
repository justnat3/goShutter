package main

import (
	"fmt"
	"strconv"
)

func main() {
	dir, debugging := Cli()
	// Read IOReadDir Comment

	fileName, filePath, dupespath, progress := IOReadDir(dir)
	//Get Items already in dupes folder. If Dupes does not exist we can create it.
	itemsInDupes := IOReadDupeFolder(dupespath)
	///println(fileName, filePath, dupespath, progress)
	// Read Hashing Comment
	logs := HashFiles(fileName, filePath, dupespath, progress)
	itemsCaught := IOReadDupeFolder(dupespath)

	if debugging {
		fmt.Println("DEBUG:")
		fmt.Printf("\n-------------------------------------------------------------------------------------\n")
		for i := 0; i < len(logs); i++ {
			fmt.Printf("Failed to open: %s\n", logs[i])
		}
		fmt.Printf("\n-------------------------------------------------------------------------------------\n")
	}

	//We can assume that the items-caught were not already in the dupes folder so -> substract the total
	fmt.Printf("Path to Dupes: %s\n", dupespath)
	fmt.Printf("Items Caught: %s\n", strconv.Itoa(itemsCaught-itemsInDupes))

	// Do not escape if run from exe
	fmt.Scanln("enter: ")
}

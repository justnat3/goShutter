package main

import (
	"fmt"
	"strconv"
)

func main() {
	dir := Cli()
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

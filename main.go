//This is a main.go file. Treat it well. Treat it with kindness. Don't forget to struggle.
package main

import (
	//"encoding/hex"
	"fmt"
	//"log"
	//"github.com/devedge/imagehash"

	"io/ioutil"
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

	window := app.NewWindow()
	for e := range window.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			// The window was closed.
			return e.Err
		case system.FrameEvent:
			// A request to draw the window state.
			ops := new(op.Ops)
			// Draw the state into ops.
			draw(ops)
			// Update the display.
			e.Frame(ops)
		}
	}

	files, err := IOReadDir("C:\\git\\test_dir")

	if err != nil {
		println("Did not Preform IOREAD_DIR Function")
	}

	time.Sleep(5 * time.Second)
	IOReadFile(files)
}

//IOReadFile : Take in files from IOREADDir function and read the bytes to check contentType
func IOReadFile(files []string) {
	var read_file string
	n := 512

	buf := make([]byte, n)

	for i := 0; i < len(files); i++ {
		fmt.Println("FOR=> " + files[i])

		read_file = files[i]
		read, err := ioutil.ReadFile(read_file)

		println(read)
		if err != nil {
			println(err)
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

	println(root + "\\")
	for _, file := range fileInfo {
		c++
		file_path := root + "\\" + file.Name()
		files = append(files, file_path)
	}
	return files, nil
}

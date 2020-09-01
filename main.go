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

type hash struct {
	path string
	hash string
}

func main() {
	files, dupespath, err := IOReadDir("C:\\Users\\Nathan Reed\\Downloads\\afreightdata\\afreightdata" + "\\")

	if err != nil {
		panic(err)
	}

	HashFiles(files, dupespath)
	fmt.Scanln("enter: ")
}

//HashFiles : take array of files and hash them
func HashFiles(files []string, dupespath string) {

	//TODO:
	// Grab File name and append it to the end of the dupe_fileName
	// hashes are the key in the hash table
	// mapping members will not take place if a member exists in the hashtable

	var readFile string
	const BlockSize = 64

	m := make(map[string]string)

	for i := 0; i < len(files); i++ {

		readFile = files[i]
		f, err := os.Open(readFile)
		defer f.Close()
		println(f.Name())
		if err != nil {
			log.Fatal("Could not open file")
		}

		buff := make([]byte, 512)
		f.Read(buff)
		newFile := dupespath + f.Name()
		println(newFile)
		println(dupespath)
		switch {
		case http.DetectContentType(buff) == "image/jpeg":
			hasher := sha256.New()

			if _, err := io.Copy(hasher, f); err != nil {
				log.Fatal(err)
			}
			sum := hasher.Sum(nil)

			_, ok := m[hex.EncodeToString(sum)]
			if ok == false {
				continue
			} else if ok == true {
				CopyFile(readFile, newFile)
			}

			m[hex.EncodeToString(sum)] = readFile
			f.Close()
		case http.DetectContentType(buff) == "image/png":
			hasher := sha256.New()

			if _, err := io.Copy(hasher, f); err != nil {
				log.Fatal(err)
			}

			sum := hasher.Sum(nil)

			_, ok := m[hex.EncodeToString(sum)]
			if ok == false {
				println("\nFalse\n")
			} else {
				CopyFile(readFile, newFile)
				fmt.Printf("SOURCE =>" + readFile + "\n")
				fmt.Printf("\nDEST =>" + newFile + "\n")
			}
			m[hex.EncodeToString(sum)] = readFile
			f.Close()
		default:
			f.Close()
		}
	}
}

//IOReadDir : Read in Directory and spit out file names + PATH
func IOReadDir(root string) ([]string, string, error) {

	var files []string
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

	println("Scanning...  " + root + "/\n")

	for _, file := range fileInfo {
		c++
		filePath := root + file.Name()
		files = append(files, filePath)
	}
	return files, dupespath, nil
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

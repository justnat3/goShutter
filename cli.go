package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thatisuday/commando"
)

var dir string

// Cli : add cli functionality to goDupe
func Cli() string {
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

	return dir

}

package main

import (
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	//var for number of arguments
	argsNum := len((os.Args)) - 1

	//if no args given
	if argsNum == 0 {
		fmt.Println("EpubGoTerminal: \n\tGo utility to print EPUB chapters into the terminal.")
		fmt.Println("Use:")
		fmt.Println("\tepubgot epubfile\t\t# list chapters of the EPUB and print image count")
		fmt.Println("\tepubgot epubfile chapternumber\t# print a chapter in the EPUB")
		fmt.Println("\tepubgot epubfile all\t\t# print all chapters of the EPUB")
		fmt.Println("Author:")
		fmt.Println("\tAhmed (github.com/ahmeddots)")
		os.Exit(0)
	}

	//recovery function
	defer printErrorAndExit()

	//check if epub
	if filepath.Ext(os.Args[1]) != ".epub" {
		panic(fmt.Errorf("(1) file is not an epub: %v", os.Args[1]))
	}

	//open epub
	archv, err := zip.OpenReader(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("(1) failed to open epub: %v", os.Args[1]))
	}
	defer archv.Close()

	switch argsNum {
	//LIST when given one argument
	case 1:
		listChapters(archv)

	//OPEN when given two arguments
	case 2:
		if os.Args[2] == "all" {
			printAllChapters(archv)
			os.Exit(0)
		}

		chapterNum, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(fmt.Errorf(("(1) not given a number for wanted chapter")))
		}
		printChapter(archv, chapterNum)
	}
}

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return file
}

func main() {
	targetFile := flag.String("file", "", "")
	old := flag.String("old", "", "old string that needs to be changed")
	new := flag.String("new", "", "new string that will replace the old one")
	all := flag.Bool("all", true, "if set to true, it will replace all the entries in a file. If set to false, will only replace the first one")
	verbose := flag.Bool("v", false, "whetever to be verbose or not")
	flag.Parse()
	if *verbose == true {
		fmt.Println("verbosity set to true")
		fmt.Println("reading file " + *targetFile)
		fmt.Println("old string " + *old)
		fmt.Println("new string " + *new)
		fmt.Println("replace all entries? (if set to false, will only replace the first one) " + strconv.FormatBool(*all))
		fmt.Println("")
	}
	fileContents := string(readFile(*targetFile))
	var newContent []byte
	if *all == false {
		newContent = []byte(strings.Replace(fileContents, *old, *new, 1))
	}
	if *all == true {
		newContent = []byte(strings.ReplaceAll(fileContents, *old, *new))
	}
	err := ioutil.WriteFile(*targetFile, newContent, 0777)
	if err != nil {
		panic(err)
	}
	fmt.Println("Wrote to " + *targetFile + " correctly. Exiting...")
}

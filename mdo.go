package main

import (
	"flag"
	"fmt"
	"mongo_data_object/internal"
)

// main
func main() {

	var goptions = internal.MDOOptions{}
	var help bool

	flag.BoolVar(&help, "help", false, "Print help message")
	flag.BoolVar(&goptions.Rebuild, "rebuild", false, "rebuild all files")

	flag.Parse()

	if help {
		fmt.Println(internal.Name)
		fmt.Printf("MDO Processing: Version: %s\n", internal.Version)
		flag.PrintDefaults()
		return
	}

	internal.MDOProcess(&goptions)

	fmt.Println("")

}

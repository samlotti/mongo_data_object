package internal

import "fmt"

var Version = "0.1.0"
var Name = "Mongo Data Objects"

type MDOOptions struct {
	Sdir              string
	Rebuild           bool
	Watch             bool
	SupportBranch     string
	RenderLineNumbers bool
}

func GteProcess(opt *MDOOptions) {

	fmt.Printf(" __          __     ___  ___        __            ___  ___  __  \n|__) |    | |__)     |  |__   |\\/| |__) |     /\\   |  |__  /__` \n|__) |___ | |        |  |___  |  | |    |___ /~~\\  |  |___ .__/ \n")

	fmt.Printf("Blip Processing: Version: %s\n", Version)
	fmt.Printf("Rebuild All: %v\n", opt.Rebuild)
	fmt.Printf("Render: LineNumbers %v\n", opt.Rebuild)
	fmt.Printf("Source folder: %s\n", opt.Sdir)

	//processDir(opt.Sdir, opt)
	//fmt.Printf("\n")
	//
	//if opt.Watch {
	//	watchFiles(opt)
	//}
}

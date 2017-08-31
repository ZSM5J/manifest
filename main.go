package main

import (
	"github.com/ZSM5J/config"
	"github.com/ZSM5J/manifest"
	"flag"
	"fmt"
)

var (
	option = flag.String("o", "", "Choose option: start/migrate")
)

func main() {
	flag.Parse()

	config.LoadConfiguration("./config.json")

	//Chose start option
	switch op := *option; op {
	case "read":
		manifest.ReadFiles(config.Config.Folders)
	case "merge":
		fmt.Println("merge")
		manifest.SnapshotList(1)
		manifest.PromptCycle()
	default:
		fmt.Println("You need to choose option")
	}
}




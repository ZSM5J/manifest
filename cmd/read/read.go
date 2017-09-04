package main

import (
	"manifest/config"
	"manifest/manifest"
	"fmt"
)



func main() {


	config.LoadConfiguration("./config.json")
	fmt.Println(config.Config)
	manifest.ReadFiles(config.Config.Folders)


}




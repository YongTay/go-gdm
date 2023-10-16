package main

import (
	"flag"
	"fmt"
	gdm "go-gdm/core"
)

const Version = "v0.0.1-beta"

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Println(Version)
		return
	}
	gdm.Run()
}

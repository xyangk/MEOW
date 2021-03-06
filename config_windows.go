package main

import (
	"os"
	"path"
)

const (
	rcFname           = "rc.txt"
	alwaysDirectFname = "direct.txt"

	newLine = "\r\n"
)

func initConfigDir() {
	// On windows, put the configuration file in the same directory of meow executable
	// This is not a reliable way to detect binary directory, but it works for double click and run
	configPath.dir = path.Dir(os.Args[0])
}

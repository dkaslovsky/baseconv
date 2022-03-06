package main

import (
	"fmt"
	"os"

	"github.com/dkaslovsky/baseconv/cmd"
)

const (
	name    = "baseconv"
	version = "0.0.1"
)

func main() {
	err := cmd.Run(name, version, os.Args)
	if err != nil {
		fmt.Printf("%s: %v\n", name, err)
		os.Exit(1)
	}
}

package main

import (
	"log"
	"os"

	"github.com/dkaslovsky/baseconv/cmd"
)

func main() {
	err := cmd.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

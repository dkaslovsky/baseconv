package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dkaslovsky/baseconv/cmd/decode"
	"github.com/dkaslovsky/baseconv/cmd/encode"
)

func main() {
	if len(os.Args) == 1 {
		usage()
		return
	}

	var err error

	switch os.Args[1] {
	case "encode":
		err = encode.Run(os.Args[2:])
	case "decode":
		err = decode.Run(os.Args[2:])
	default:
		log.Fatalf("received unknown command %s", os.Args[1])
	}

	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Println("usage goes here")
}

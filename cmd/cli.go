package cmd

import (
	"fmt"

	"github.com/dkaslovsky/baseconv/cmd/decode"
	"github.com/dkaslovsky/baseconv/cmd/encode"
)

// Run executes the top level command
func Run(cliArgs []string) error {

	if len(cliArgs) <= 1 {
		printUsage()
		return nil
	}

	cmd, args := cliArgs[1], cliArgs[2:]

	switch cmd {
	case "encode":
		return encode.Run(args)
	case "decode":
		return decode.Run(args)
	case "help", "--help", "-h":
		printUsage()
		return nil
	default:
		return fmt.Errorf("received unknown command %s", cmd)
	}

}

func printUsage() {
	fmt.Println("usage goes here")
}

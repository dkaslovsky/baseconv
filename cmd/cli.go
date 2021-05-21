package cmd

import (
	"fmt"

	"github.com/dkaslovsky/baseconv/cmd/decode"
	"github.com/dkaslovsky/baseconv/cmd/encode"
)

const name = "baseconv"

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
	fmt.Printf("%s converts between base 10 integers and string representations in arbitraty bases\n\n", name)
	fmt.Print("Usage:\n")
	fmt.Printf("  %s command [flags]\n\n", name)
	fmt.Print("Commands:\n")
	fmt.Print("  encode - encodes a base 10 integer in a new base\n")
	fmt.Print("  decode - decodes a string representation of a base 10 integer\n")
}

package cmd

import (
	"fmt"

	"github.com/dkaslovsky/baseconv/cmd/decode"
	"github.com/dkaslovsky/baseconv/cmd/encode"
)

// Run executes the top level command
func Run(name string, version string, cliArgs []string) error {
	if len(cliArgs) <= 1 {
		printUsage(name)
		return nil
	}

	cmd, args := cliArgs[1], cliArgs[2:]

	switch cmd {
	case "encode":
		return encode.Run(args)
	case "decode":
		return decode.Run(args)
	case "-help", "-h":
		printUsage(name)
		return nil
	case "-version", "-v":
		printVersion(name, version)
		return nil
	default:
		return fmt.Errorf("unknown command %s", cmd)
	}
}

func printUsage(name string) {
	fmt.Printf("%s converts between base 10 integers and string representations in arbitraty bases\n", name)

	fmt.Print("\nUsage:\n")
	fmt.Printf("  %s [flags]\n", name)
	fmt.Printf("  %s [command]\n", name)

	fmt.Print("\nAvailable Commands:\n")
	fmt.Print("  encode\tencodes a base 10 integer in a new base\n")
	fmt.Print("  decode\tdecodes a string representation of a base 10 integer\n")

	fmt.Print("\nFlags:\n")
	fmt.Printf("  -h, -help\thelp for %s\n", name)
	fmt.Printf("  -v, -version\tversion for %s\n", name)
}

func printVersion(name string, version string) {
	fmt.Printf("%s v%s\n", name, version)
}

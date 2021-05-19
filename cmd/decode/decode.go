package decode

import (
	"errors"
	"flag"
	"fmt"

	"github.com/dkaslovsky/baseconv/pkg/alphabet"
	b "github.com/dkaslovsky/baseconv/pkg/base"
)

type cmdOpts struct {
	// command flags
	base uint64
	// positional args
	enc string
}

// Run executes the decode (sub)command
func Run(args []string) error {
	cmd := flag.NewFlagSet("decode", flag.ExitOnError)
	opts := &cmdOpts{}
	attachOpts(cmd, opts)
	setUsage(cmd)

	err := parseArgs(cmd, opts, args)
	if err == errNoArgs {
		cmd.Usage()
		return nil
	}
	if err != nil {
		return err
	}

	numeric, err := alphabet.FromString(opts.enc)
	if err != nil {
		return err
	}

	dec, err := b.ToBase10(numeric, opts.base)
	if err != nil {
		return err
	}

	fmt.Println(dec)
	return nil
}

func attachOpts(cmd *flag.FlagSet, opts *cmdOpts) {
	cmd.Uint64Var(&opts.base, "b", 0, "new base to encode input integer")
	cmd.Uint64Var(&opts.base, "base", 0, "new base to encode input integer")
}

// errorNoArgs is returned when no arguments are passed to the command
var errNoArgs = errors.New("missing required argument(s)")

func parseArgs(cmd *flag.FlagSet, opts *cmdOpts, args []string) error {
	if len(args) == 0 {
		return errNoArgs
	}

	err := cmd.Parse(args)
	if err != nil {
		return err
	}

	// handle positional argument(s)
	if cmd.NArg() != 1 {
		return errors.New("must specify encoded string as single positional argument")
	}
	opts.enc = cmd.Arg(0)

	return nil
}

func setUsage(cmd *flag.FlagSet) {
	cmd.Usage = func() {
		fmt.Printf("%s decodes a representation of a base 10 integer\n\n", cmd.Name())
		fmt.Printf("%s [flags] stringRep\n\n", cmd.Name())
		fmt.Printf("Args:\n  stringRep - string representation of an encoded base 10 integer to decode (required)\n\n")
		fmt.Printf("Flags:\n")
		cmd.PrintDefaults()
	}
}

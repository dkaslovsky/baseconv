package decode

import (
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
	opts, err := parseArgs(cmd, args)
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

func parseArgs(cmd *flag.FlagSet, args []string) (*cmdOpts, error) {
	opts := &cmdOpts{}
	attachOpts(cmd, opts)

	err := cmd.Parse(args)
	if err != nil {
		return opts, err
	}

	// handle positional argument(s)
	if cmd.NArg() != 1 {
		return opts, fmt.Errorf("must specify encoded string as single positional argument")
	}
	opts.enc = cmd.Arg(0)

	return opts, nil
}

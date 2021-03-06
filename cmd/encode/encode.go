package encode

import (
	"errors"
	"flag"
	"fmt"
	"strconv"

	"github.com/dkaslovsky/baseconv/pkg/alphabet"
	"github.com/dkaslovsky/baseconv/pkg/baseconv"
)

// Run executes the encode (sub)command
func Run(args []string) error {
	cmd := flag.NewFlagSet("encode", flag.ExitOnError)
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

	return run(opts)
}

func run(opts *cmdOpts) error {
	enc, err := baseconv.FromBase10(opts.num, opts.base)
	if err != nil {
		return err
	}

	str, serr := alphabet.ToString(enc)
	if serr != nil {
		return serr
	}

	if opts.pad {
		str, err = alphabet.Pad(str, int(opts.maxDigits))
		if err != nil {
			return err
		}
	}

	fmt.Println(str)
	return nil
}

// errorNoArgs is returned when no arguments are passed to the command
var errNoArgs = errors.New("missing required argument(s)")

type cmdOpts struct {
	// command flags
	base      uint64
	maxDigits uint64
	pad       bool

	// positional args
	num uint64
}

func attachOpts(cmd *flag.FlagSet, opts *cmdOpts) {
	cmd.Uint64Var(&opts.base, "b", 0, "new base to encode input integer")
	cmd.Uint64Var(&opts.base, "base", 0, "new base to encode input integer")

	cmd.Uint64Var(&opts.maxDigits, "d", 0, "maximum number of digits to use for encoding")
	cmd.Uint64Var(&opts.maxDigits, "digits", 0, "maximum number of digits to use for encoding")

	cmd.BoolVar(&opts.pad, "p", false, "pad output to have exactly the number of specified digits")
	cmd.BoolVar(&opts.pad, "pad", false, "pad output to have exactly the number of specified digits")
}

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
		return errors.New("must specify base 10 integer to encode as single positional argument")
	}
	num, nerr := strconv.ParseUint(cmd.Arg(0), 10, 64)
	if nerr != nil {
		return fmt.Errorf("could not parse positional argument %s as uint64", cmd.Arg(0))
	}
	opts.num = num

	return validateOpts(opts)
}

func validateOpts(opts *cmdOpts) error {
	maxBase := alphabet.Len()
	if opts.base > maxBase {
		return fmt.Errorf("base [%d] exceeds alphabet size [%d]", opts.base, maxBase)
	}
	maxNum, err := baseconv.GetLargestBase10(opts.base, opts.maxDigits)
	if err != nil {
		return err
	}
	if opts.num > maxNum {
		return fmt.Errorf("cannot encode %d in base %d with %d digits", opts.num, opts.base, opts.maxDigits)
	}
	return nil
}

func setUsage(cmd *flag.FlagSet) {
	cmd.Usage = func() {
		fmt.Printf("%s encodes a base 10 integer in a new base\n\n", cmd.Name())

		fmt.Print("Usage:\n")
		fmt.Printf("  %s [flags] NUM\n\n", cmd.Name())

		fmt.Print("Args:\n")
		fmt.Printf("  NUM\tpositive base 10 integer to encode (required)\n\n")

		fmt.Printf("Flags:\n")
		cmd.PrintDefaults()
	}
}

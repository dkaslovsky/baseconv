package encode

import (
	"fmt"

	"github.com/dkaslovsky/baseconv/pkg/alphabet"
	b "github.com/dkaslovsky/baseconv/pkg/base"
)

// these will be cmdline args/flags
var (
	base      uint64 = 62
	maxDigits uint64 = 3
	num       uint64 = 1000
	pad       bool   = true
)

// Run executes the encode (sub)command
func Run(args []string) error {
	// validate input
	maxNum, err := b.GetLargestBase10Number(base, maxDigits)
	if err != nil {
		return err
	}
	if num > maxNum {
		return fmt.Errorf("cannot encode base 10 number %d with %d digits", num, maxDigits)
	}

	enc, err := b.FromBase10(num, base)
	if err != nil {
		return err
	}

	if pad {
		str, err := alphabet.ToPaddedString(enc, int(maxDigits))
		if err != nil {
			return err
		}
		fmt.Println(str)
		return nil
	}

	str, err := alphabet.ToString(enc)
	if err != nil {
		return err
	}
	fmt.Println(str)
	return nil
}

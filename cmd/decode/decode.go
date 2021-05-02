package decode

import (
	"fmt"

	"github.com/dkaslovsky/baseconv/pkg/alphabet"
	b "github.com/dkaslovsky/baseconv/pkg/base"
)

// these will be cmdline args/flags
var (
	base uint64 = 62
	str  string = "0g8"
)

// Run executes the decode (sub)command
func Run(args []string) error {
	numeric, err := alphabet.FromString(str)
	if err != nil {
		return err
	}

	dec, err := b.ToBase10(numeric, base)
	if err != nil {
		return err
	}

	fmt.Println(dec)
	return nil
}

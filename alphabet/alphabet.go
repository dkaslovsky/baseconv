package alphabet

import (
	"fmt"
	"strings"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Size is the length of the alphabet
var Size = len(alphabet)

// ToString maps a slice of numbers to a string of corresponding characters
func ToString(numeric []uint64) (string, error) {
	sz := uint64(Size)
	str := []byte{}
	for _, n := range numeric {
		if n >= sz {
			return "", fmt.Errorf("value %d exceeds alphabet size of %d", n, sz)
		}
		str = append(str, alphabet[n])
	}
	return string(str), nil
}

// FromString maps a string of characters to a slice of corresponding numbers
func FromString(str string) ([]uint64, error) {
	numeric := []uint64{}
	for _, s := range str {
		i := strings.Index(alphabet, string(s))
		if i == -1 {
			return numeric, fmt.Errorf("character %c not found in alphabet", s)
		}
		numeric = append(numeric, uint64(i))
	}
	return numeric, nil
}

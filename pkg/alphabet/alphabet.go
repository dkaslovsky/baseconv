package alphabet

import (
	"fmt"
	"strings"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// FromString maps a string of characters to a slice of corresponding numbers
func FromString(str string) ([]uint64, error) {
	numeric := []uint64{}
	for _, s := range str {
		i := strings.Index(alphabet, string(s))
		if i == -1 {
			return numeric, fmt.Errorf("character [%c] not found in alphabet", s)
		}
		numeric = append(numeric, uint64(i))
	}
	return numeric, nil
}

// ToString maps a slice of numbers to a string of corresponding characters
func ToString(numeric []uint64) (string, error) {
	str := []byte{}
	for _, n := range numeric {
		if n >= Len() {
			return "", fmt.Errorf("value [%d] cannot be represented in alphabet of size [%d]", n, Len())
		}
		str = append(str, alphabet[n])
	}
	return string(str), nil
}

// Pad appends the zero character of the alphabet to a string to produce a string of desired length
func Pad(str string, strLen int) (string, error) {
	padLen := strLen - len(str)
	if padLen < 0 {
		return "", fmt.Errorf("input string length [%d] exceeds desired padded length [%d]", len(str), strLen)
	}
	padding := strings.Repeat(Zero(), padLen)
	return padding + str, nil
}

// Len returns the length of the alphabet
func Len() uint64 {
	return uint64(len(alphabet))
}

// Zero returns the alphabet's index-zero character used for padding a string
func Zero() string {
	return string(alphabet[0])
}

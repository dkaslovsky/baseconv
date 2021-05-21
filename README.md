# baseconv

baseconv is a lightweight, dependency-free, Golang package and CLI for converting between base 10 integers and string representations in an arbitrary base.

A practical use for such a package is efficiently encoding very large intergers using fewer bits than required in base 10 and, equivalently, 
converting a very large number of intergers (e.g., auto-incrementing primay keys) to short, unique strings as is done by URL shorterners.

For example, the `1,000,000,000,000`th URL stored in a database is efficiently mapped to the unique slug `hBxM5A4`.  Given this slug as part of a URL (e.g., `myshorturl.com/hBxM5A4`), it can be uniquely mapped back to the integer that serves as the primary key in a database storing the full URL to which the user is to be redirected.

baseconv uses an alphabet that supports encoding bases between 2 and 62, inclusive.

## CLI Usage

To use baseconv as a CLI, call either the `encode` or `decode` commands as described by the CLI's help:
```
$ baseconv -h

baseconv converts between base 10 integers and string representations in arbitraty bases

Usage:
  baseconv command [flags]

Commands:
  encode - encodes a base 10 integer in a new base
  decode - decodes a string representation of a base 10 integer
```
The `encode` command accepts a base 10 integer as its only positional argument and flags specify the new base, the maximum number of digits to be used, and whether the result should be padded to contain exactly that number of digits:
```
$ baseconv encode -h

encode encodes a base 10 integer in a new base

Usage:
  encode [flags] base10Int

Args:
  base10Int - positive base 10 integer to encode (required)

Flags:
  -b uint
    	new base to encode input integer
  -base uint
    	new base to encode input integer
  -d uint
    	maximum number of digits to use for encoding
  -digits uint
    	maximum number of digits to use for encoding
  -p	pad output to have exactly the number of specified digits
  -pad
    	pad output to have exactly the number of specified digits
```

For example,
```
$ baseconv encode -b 62 -d 7 1000000000000
hBxM5A4
```

The `decode` command performs the inverse of the encoding: it converts a string representation in a specified base to a base 10 integer.  It accepts the string as its only positional argument and the base is specified as a flag:
```
$ baseconv decode -h

decode decodes a string representation of a base 10 integer from an arbitrary base

Usage:
  decode [flags] stringRep

Args:
  stringRep - string representation of an encoded base 10 integer to decode (required)

Flags:
  -b uint
    	new base to encode input integer
  -base uint
    	new base to encode input integer
```

For example,
```
$ baseconv decode -b 62 hBxM5A4
1000000000000
```

## Package Usage

baseconv provides two packages that can be imported for use in other projects:
- the `baseconv` package implements the conversion between bases
- the `alphabet` package implements the conversion between numeric arrays and string representations.

### baseconv
The `baseconv` package is imported as
```go
import "github.com/dkaslovsky/baseconv/pkg/baseconv"
```
and provides the following functions:
```go
// FromBase10 converts a base 10 number to a slice representing the number in a specified base
func FromBase10(num, base uint64) ([]uint64, error)

// ToBase10 converts a number in a specified base represented by a slice into its base 10 value
func ToBase10(num []uint64, base uint64) (uint64, error)

// GetLargestBase10Number returns the largest base 10 number that can be represented
// in the specified base with the specified number of digits
func GetLargestBase10Number(base, digits uint64) (uint64, error)
```

### alphabet
The `alphabet` package is imported as
```go
import "github.com/dkaslovsky/baseconv/pkg/alphabet"
```
and provides the following functions:
```go
// FromString maps a string of characters to a slice of corresponding numbers
func FromString(str string) ([]uint64, error)

// ToString maps a slice of numbers to a string of corresponding characters
func ToString(numeric []uint64) (string, error)

// Pad appends the zero character of the alphabet to a string to produce a string of desired length
func Pad(str string, strLen int) (string, error)

// Len returns the length of the alphabet
func Len() uint64

// Zero returns the alphabet's index-zero character used for padding a string
func Zero() string
```
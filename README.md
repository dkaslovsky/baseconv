# baseconv

Lightweight, dependency-free, Golang package and CLI for converting between base 10 integers and string representations in an arbitrary base.

## Overview

baseconv can be used to efficiently encode very large integers into short, unique strings ("slugs").
By converting a large base 10 integer to a different base, the original integer is encoded in a string representation using fewer bits.

URL shorteners often use such encodings to map integer database keys to unique slugs that serve as shortened URLs.
For example, the `1,000,000,000,000`th URL stored in a database can be mapped to the unique slug `hBxM5A4` by converting to base 62.  Given this slug as part of a URL (e.g., `shorturl.xyz/hBxM5A4`), it can be uniquely mapped back to the integer key used to lookup the full URL.

baseconv uses an alphabet that supports encoding in base b, 2 <= b <= 62.

## CLI Usage

To use baseconv as a CLI, call either the `encode` or `decode` commands as described by the CLI's help:
```
$ baseconv -h

baseconv converts between base 10 integers and string representations in arbitraty bases

Usage:
  baseconv [flags]
  baseconv [command]

Available Commands:
  encode	encodes a base 10 integer in a new base
  decode	decodes a string representation of a base 10 integer

Flags:
  -h, -help	help for baseconv
  -v, -version	version for baseconv
```
The `encode` command accepts a base 10 integer as its only positional argument and flags specify the new base, the maximum number of digits to be used, and whether the result should be padded to contain exactly that number of digits:
```
$ baseconv encode -h

encode encodes a base 10 integer in a new base

Usage:
  encode [flags] NUM

Args:
  NUM	positive base 10 integer to encode (required)

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
$ baseconv encode -b 62 -d 7 1000000000001
hBxM5A5
```

The `decode` command performs the inverse of the encoding: it converts a string representation in a specified base to a base 10 integer.  It accepts the string as its only positional argument and the base is specified as a flag:
```
$ baseconv decode -h

decode decodes a string representation of a base 10 integer from an arbitrary base

Usage:
  decode [flags] STRINGREP

Args:
  STRINGREP	string representation of an encoded base 10 integer to decode (required)

Flags:
  -b uint
    	base of input number
  -base uint
    	base of input number
```

For example,
```
$ baseconv decode -b 62 hBxM5A5
1000000000001
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
func FromBase10(num uint64, base uint64) ([]uint64, error)

// ToBase10 converts a number in a specified base represented by a slice into its base 10 value
func ToBase10(num []uint64, base uint64) (uint64, error)

// GetLargestBase10 returns the largest base 10 number that can be represented
// in the specified base with the specified number of digits
func GetLargestBase10(base uint64, digits uint64) (uint64, error)
```

### alphabet
The `alphabet` package is imported as
```go
import "github.com/dkaslovsky/baseconv/pkg/alphabet"
```
and provides the following functions:
```go
// FromString converts a string of characters to a slice of corresponding numbers
func FromString(str string) ([]uint64, error)

// ToString converts a slice of numbers to a string of corresponding characters
func ToString(numeric []uint64) (string, error)

// Pad appends the zero character of the alphabet to a string to produce a string of desired length
func Pad(str string, strLen int) (string, error)

// Len returns the length of the alphabet
func Len() uint64

// Zero returns the alphabet's index-zero character used for padding a string
func Zero() string
```

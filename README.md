# baseconv

baseconv is a lightweight, dependency-free, Golang package and CLI for converting between base 10 integers and string representations in an arbitrary base.

A practical use for such a package is efficiently encoding very large intergers using fewer bits than required in base 10 and, equivalently, 
converting a very large number of intergers (e.g., auto-incrementing primay keys) to short, unique strings as is done by URL shorterners.

For example, the `1,000,000,000,000`th URL stored in a database is efficiently mapped to the unique slug `hBxM5A4`.  Given this slug as part of a URL (e.g., `myshorturl.com/hBxM5A4`), it can be uniquely mapped back to the integer that serves as the primary key in a database storing the full URL to which the user is to be redirected.

baseconv uses an alphabet that supports encoding is bases between 2 and 62, inclusive.

## Usage

### CLI

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

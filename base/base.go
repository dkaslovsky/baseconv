package base

import (
	"fmt"
	"math"
)

// FromBase10 converts a base 10 number to a slice representing the number in a specified base
func FromBase10(num, base int64) ([]int64, error) {
	if err := validateBase(base); err != nil {
		return []int64{0}, err
	}

	if num == 0 {
		return []int64{0}, nil
	}

	b := float64(base)
	numDigits := int64(math.Floor(1.0 + math.Log(float64(num))/math.Log(b)))
	newBaseDigits := make([]int64, numDigits)

	var i int64
	for i = 0; i < numDigits; i++ {
		exponent := numDigits - i - 1
		digitInBase10 := int64(math.Pow(b, float64(exponent)))
		newBaseDigits[i] = num / digitInBase10 // integer division
		num = num % digitInBase10
	}

	return newBaseDigits, nil
}

// ToBase10 converts a number in a specified base represented by a slice into its base 10 value
func ToBase10(num []int64, base int64) (int64, error) {
	if err := validateBase(base); err != nil {
		return 0, err
	}

	b := float64(base)
	numDigits := len(num)

	var base10 int64 = 0

	for i, n := range num {
		if n >= base {
			return 0, fmt.Errorf("digit %d must be less than base %d", n, base)
		}
		exponent := numDigits - i - 1
		base10 += n * int64(math.Pow(b, float64(exponent)))
	}
	return base10, nil
}

func validateBase(base int64) error {
	if base < 2 {
		return fmt.Errorf("base cannot be less than 2")
	}
	return nil
}

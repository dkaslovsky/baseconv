package baseconv

import (
	"fmt"
	"math"
)

// roundoffTol is the tolerance for detecting roundoff error
const roundoffTol = 1e-8

// FromBase10 converts a base 10 number to a slice representing the number in a specified base
func FromBase10(num uint64, base uint64) ([]uint64, error) {
	if err := validateBase(base); err != nil {
		return nil, err
	}
	if num == 0 {
		return []uint64{0}, nil
	}

	b := float64(base)
	n := float64(num)
	numDigits := getNumDigits(n, b, roundoffTol)
	newBaseDigits := make([]uint64, numDigits)

	for i := 0; i < numDigits; i++ {
		exponent := numDigits - i - 1
		digitInBase10 := uint64(math.Pow(b, float64(exponent)))
		newBaseDigits[i] = num / digitInBase10 // integer division
		num = num % digitInBase10
	}

	return newBaseDigits, nil
}

// ToBase10 converts a number in a specified base represented by a slice into its base 10 value
func ToBase10(num []uint64, base uint64) (uint64, error) {
	if err := validateBase(base); err != nil {
		return 0, err
	}

	b := float64(base)
	numDigits := len(num)
	base10 := uint64(0)

	for i, n := range num {
		if n >= base {
			return 0, fmt.Errorf("cannot convert digit [%d] to base [%d]", n, base)
		}
		exponent := numDigits - i - 1
		base10 += n * uint64(math.Pow(b, float64(exponent)))
	}

	return base10, nil
}

// GetLargestBase10 returns the largest base 10 number that can be represented
// in the specified base with the specified number of digits
func GetLargestBase10(base uint64, digits uint64) (uint64, error) {
	if err := validateBase(base); err != nil {
		return 0, err
	}
	return uint64(math.Pow(float64(base), float64(digits)) - 1), nil
}

func validateBase(base uint64) error {
	if base < 2 {
		return fmt.Errorf("base cannot be less than 2")
	}
	return nil
}

func getNumDigits(num float64, base float64, tol float64) int {
	logN := math.Log(num) / math.Log(base)

	// correct for roundoff error near an integer value of logN that causes problems with flooring;
	// this usually happens when num is an exact power of base
	rounded := math.Round(logN)
	if math.Abs(rounded-logN) < tol {
		return int(rounded) + 1
	}

	return int(math.Floor(logN)) + 1
}

package baseconv

import (
	"math"
	"testing"
)

func TestFromBase10(t *testing.T) {
	type testCase struct {
		num      uint64
		base     uint64
		expected []uint64
	}

	tests := map[string]testCase{
		"convert zero": {
			num:      0,
			base:     2,
			expected: []uint64{0},
		},
		"convert to binary": {
			num:      5,
			base:     2,
			expected: []uint64{1, 0, 1},
		},
		"convert to binary with trailing zeros": {
			num:      12,
			base:     2,
			expected: []uint64{1, 1, 0, 0},
		},
		"convert to base 62 with number less than 62": {
			num:      5,
			base:     62,
			expected: []uint64{5},
		},
		"convert to base 62 with number equal to 62": {
			num:      62,
			base:     62,
			expected: []uint64{1, 0},
		},
		"convert to base 62 with number larger than 62": {
			num:      3_520_000_000_000,
			base:     62,
			expected: []uint64{61, 60, 14, 45, 17, 10, 24},
		},
		"convert to base 10 from base 10": {
			num:      587,
			base:     10,
			expected: []uint64{5, 8, 7},
		},
		"convert to base 10 from base 10 with multiple of 10": {
			num:      1000,
			base:     10,
			expected: []uint64{1, 0, 0, 0},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			res, err := FromBase10(test.num, test.base)
			if err != nil {
				t.Fatalf("unexpected non nil error: %v", err)
			}
			if len(res) != len(test.expected) {
				t.Errorf("result %v not equal to expected %v", res, test.expected)
				return
			}
			for i := 0; i < len(res); i++ {
				if res[i] != test.expected[i] {
					t.Errorf("result %v not equal to expected %v", res, test.expected)
					return
				}
			}
		})
	}
}

func TestFromBase10WithError(t *testing.T) {
	type testCase struct {
		num  uint64
		base uint64
	}

	tests := map[string]testCase{
		"target base 0": {
			num:  1,
			base: 0,
		},
		"target base 1": {
			num:  0,
			base: 1,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			test := test
			_, err := FromBase10(test.num, test.base)
			if err == nil {
				t.Fatal("expected non nil error")
			}
		})
	}
}

func TestToBase10(t *testing.T) {
	type testCase struct {
		num      []uint64
		base     uint64
		expected uint64
	}

	tests := map[string]testCase{
		"convert single zero": {
			num:      []uint64{0},
			base:     2,
			expected: 0,
		},
		"convert multiple zeros": {
			num:      []uint64{0, 0, 0, 0, 0, 0},
			base:     2,
			expected: 0,
		},
		"convert from binary": {
			num:      []uint64{1, 0, 1},
			base:     2,
			expected: 5,
		},
		"convert from binary with leading zeros": {
			num:      []uint64{0, 0, 1, 1},
			base:     2,
			expected: 3,
		},
		"convert from binary with trailing zeros": {
			num:      []uint64{1, 1, 0, 0},
			base:     2,
			expected: 12,
		},
		"convert from base 62 with number less than 62": {
			num:      []uint64{5},
			base:     62,
			expected: 5,
		},
		"convert from base 62 with number equal to 62": {
			num:      []uint64{1, 0},
			base:     62,
			expected: 62,
		},
		"convert from base 62 number with number larger than 62": {
			num:      []uint64{61, 60, 14, 45, 17, 10, 24},
			base:     62,
			expected: 3_520_000_000_000,
		},
		"convert from base 10 to base 10": {
			num:      []uint64{5, 8, 7},
			base:     10,
			expected: 587,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			res, err := ToBase10(test.num, test.base)
			if err != nil {
				t.Fatalf("unexpected non nil error: %v", err)
			}
			if res != test.expected {
				t.Errorf("result %d not equal to expected %d", res, test.expected)
			}
		})
	}
}

func TestToBase10WithError(t *testing.T) {
	type testCase struct {
		num  []uint64
		base uint64
	}

	tests := map[string]testCase{
		"convert number equal to target base": {
			num:  []uint64{5},
			base: 5,
		},
		"convert number larger than target base": {
			num:  []uint64{501},
			base: 500,
		},
		"target base 0": {
			num:  []uint64{1},
			base: 0,
		},
		"target base 1": {
			num:  []uint64{0},
			base: 1,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			_, err := ToBase10(test.num, test.base)
			if err == nil {
				t.Fatal("expected non nil error")
			}
		})
	}
}

func TestGetLargestBase10(t *testing.T) {
	type testCase struct {
		base     uint64
		digits   uint64
		expected uint64
	}

	tests := map[string]testCase{
		"base=2 digits=3": {
			base:     2,
			digits:   3,
			expected: 7,
		},
		"base=2 digits=10": {
			base:     2,
			digits:   10,
			expected: 1023,
		},
		"base=10 digits=6": {
			base:     10,
			digits:   6,
			expected: 999_999,
		},
		"base=2 digits=1": {
			base:     2,
			digits:   1,
			expected: 1,
		},
		"base=200 digits=1": {
			base:     200,
			digits:   1,
			expected: 199,
		},
		"base=2 digits=0": {
			base:     2,
			digits:   0,
			expected: 0,
		},
		"base=2000 digits=0": {
			base:     2000,
			digits:   0,
			expected: 0,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			res, err := GetLargestBase10(test.base, test.digits)
			if err != nil {
				t.Fatalf("unexpected non nil error: %v", err)
			}
			if res != test.expected {
				t.Errorf("result %d not equal to expected %d", res, test.expected)
			}
		})
	}
}

func TestGetLargestBase10WithError(t *testing.T) {
	type testCase struct {
		base   uint64
		digits uint64
	}

	tests := map[string]testCase{
		"base=0 digits=10": {
			base:   0,
			digits: 10,
		},
		"base=1 digits=10": {
			base:   1,
			digits: 10,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			_, err := GetLargestBase10(test.base, test.digits)
			if err == nil {
				t.Fatal("expected non nil error")
			}
		})
	}
}

func TestGetNumDigits(t *testing.T) {
	type testCase struct {
		num      float64
		base     float64
		tol      float64
		expected int
	}

	tests := map[string]testCase{
		"exact power of 10 represented in base 10": {
			num:      1000,
			base:     10,
			tol:      1e-10,
			expected: 4,
		},
		"multiple of exact power of 10 represented in base 10": {
			num:      3000,
			base:     10,
			tol:      1e-10,
			expected: 4,
		},
		"exact power of 2 represented in base 2": {
			num:      64,
			base:     2,
			tol:      1e-10,
			expected: 7,
		},
		"one less than exact power of 2 represented in base 2": {
			num:      63,
			base:     2,
			tol:      1e-10,
			expected: 6,
		},
		"one more than exact power of 2 represented in base 2": {
			num:      65,
			base:     2,
			tol:      1e-10,
			expected: 7,
		},
		"contrived example where roundoff tolerance is used": {
			// num can't really be a float, but it is selected here to simulate the effect of roundoff
			// error in the calculation
			num:      math.Pow(2, 2.999),
			base:     2,
			tol:      1e-1,
			expected: 4,
		},
		"contrived example where roundoff tolerance is not used": {
			// num can't really be a float, but it is selected here to simulate the effect of roundoff
			// error in the calculation
			num:      math.Pow(5, 2.999),
			base:     5,
			tol:      1e-10,
			expected: 3,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			res := getNumDigits(test.num, test.base, test.tol)
			if res != test.expected {
				t.Errorf("result %d not equal to expected %d", res, test.expected)
			}
		})
	}
}

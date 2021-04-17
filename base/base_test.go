package base

import (
	"reflect"
	"testing"
)

func TestToBase10(t *testing.T) {
	type testCase struct {
		num      []int64
		base     int64
		expected int64
	}

	tests := map[string]testCase{
		"convert single zero": {
			num:      []int64{0},
			base:     2,
			expected: 0,
		},
		"convert multiple zeros": {
			num:      []int64{0, 0, 0, 0, 0, 0},
			base:     2,
			expected: 0,
		},
		"convert from binary": {
			num:      []int64{1, 0, 1},
			base:     2,
			expected: 5,
		},
		"convert from binary with leading zeros": {
			num:      []int64{0, 0, 1, 1},
			base:     2,
			expected: 3,
		},
		"convert from binary with trailing zeros": {
			num:      []int64{1, 1, 0, 0},
			base:     2,
			expected: 12,
		},
		"convert from base 62 with number less than 62": {
			num:      []int64{5},
			base:     62,
			expected: 5,
		},
		"convert from base 62 with number equal to 62": {
			num:      []int64{1, 0},
			base:     62,
			expected: 62,
		},
		"convert from base 62 number with number larger than 62": {
			num:      []int64{61, 60, 14, 45, 17, 10, 24},
			base:     62,
			expected: 3_520_000_000_000,
		},
		"convert from base 10 to base 10": {
			num:      []int64{5, 8, 7},
			base:     10,
			expected: 587,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			val, err := ToBase10(test.num, test.base)
			if err != nil {
				t.Fatalf("test case \"%s\": non nil error: %s", name, err)
			}
			if val != test.expected {
				t.Fatalf("test case \"%s\": expected %d not equal to actual %d", name, test.expected, val)
			}
		})
	}
}

func TestToBase10WithError(t *testing.T) {
	type testCase struct {
		num  []int64
		base int64
	}

	tests := map[string]testCase{
		"convert number equal to target base": {
			num:  []int64{5},
			base: 5,
		},
		"convert number larger than target base": {
			num:  []int64{501},
			base: 500,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := ToBase10(test.num, test.base)
			if err == nil {
				t.Fatalf("test case \"%s\": expected non nil error", name)
			}
		})
	}
}

func TestFromBase10(t *testing.T) {
	type testCase struct {
		num      int64
		base     int64
		expected []int64
	}

	tests := map[string]testCase{
		"convert zero": {
			num:      0,
			base:     2,
			expected: []int64{0},
		},
		"convert to binary": {
			num:      5,
			base:     2,
			expected: []int64{1, 0, 1},
		},
		"convert to binary with trailing zeros": {
			num:      12,
			base:     2,
			expected: []int64{1, 1, 0, 0},
		},
		"convert to base 62 with number less than 62": {
			num:      5,
			base:     62,
			expected: []int64{5},
		},
		"convert to base 62 with number equal to 62": {
			num:      62,
			base:     62,
			expected: []int64{1, 0},
		},
		"convert to base 62 with number larger than 62": {
			num:      3_520_000_000_000,
			base:     62,
			expected: []int64{61, 60, 14, 45, 17, 10, 24},
		},
		"convert to base 10 from base 10": {
			num:      587,
			base:     10,
			expected: []int64{5, 8, 7},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			val := FromBase10(test.num, test.base)
			if !reflect.DeepEqual(val, test.expected) {
				t.Fatalf("test case \"%s\": expected %d not equal to actual %d", name, test.expected, val)
			}
		})
	}
}

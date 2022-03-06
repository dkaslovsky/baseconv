package alphabet

import (
	"testing"
)

func TestFromString(t *testing.T) {
	type testCase struct {
		str      string
		expected []uint64
	}

	tests := map[string]testCase{
		"single zero": {
			str:      "0",
			expected: []uint64{0},
		},
		"multiple zeros": {
			str:      "0000",
			expected: []uint64{0, 0, 0, 0},
		},
		"single Z": {
			str:      "Z",
			expected: []uint64{61},
		},
		"heLLo123": {
			str:      "heLLo123",
			expected: []uint64{17, 14, 47, 47, 24, 1, 2, 3},
		},
		"heLLo123 with padding": {
			str:      "00heLLo123",
			expected: []uint64{0, 0, 17, 14, 47, 47, 24, 1, 2, 3},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			res, err := FromString(test.str)
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

func TestFromStringWithError(t *testing.T) {
	type testCase struct {
		str string
	}

	tests := map[string]testCase{
		"single value not in alphabet": {
			str: "@",
		},
		"trailing character not in alphabet": {
			str: "abc@",
		},
		"leading character not in alphabet": {
			str: "@abc",
		},
		"middle character not in alphabet": {
			str: "a@c",
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			_, err := FromString(test.str)
			if err == nil {
				t.Fatal("expected non nil error")
			}
		})
	}
}

func TestToString(t *testing.T) {
	type testCase struct {
		numeric  []uint64
		expected string
	}

	tests := map[string]testCase{
		"single zero": {
			numeric:  []uint64{0},
			expected: "0",
		},
		"multiple zeros": {
			numeric:  []uint64{0, 0, 0, 0},
			expected: "0000",
		},
		"single Z": {
			numeric:  []uint64{61},
			expected: "Z",
		},
		"heLLo123": {
			numeric:  []uint64{17, 14, 47, 47, 24, 1, 2, 3},
			expected: "heLLo123",
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			res, err := ToString(test.numeric)
			if err != nil {
				t.Fatalf("unexpected non nil error: %v", err)
			}
			if res != test.expected {
				t.Errorf("result %s not equal to expected %s", res, test.expected)
			}
		})
	}
}

func TestToStringWithError(t *testing.T) {
	type testCase struct {
		numeric []uint64
	}

	tests := map[string]testCase{
		"single value equal to alphabet size": {
			numeric: []uint64{Len()},
		},
		"single value larger than alphabet size": {
			numeric: []uint64{Len() + 1},
		},
		"mixed values of larger than and smaller than alphabet size": {
			numeric: []uint64{0, Len() + 1},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			_, err := ToString(test.numeric)
			if err == nil {
				t.Fatal("expected non nil error")
			}
		})
	}
}

func TestPad(t *testing.T) {
	type testCase struct {
		str      string
		strLen   int
		expected string
	}

	tests := map[string]testCase{
		"strLen equal to length of input": {
			str:      "heLLo",
			strLen:   5,
			expected: "heLLo",
		},
		"strLen greater than length of input": {
			str:      "heLLo",
			strLen:   10,
			expected: "00000heLLo",
		},
		"empty string input": {
			str:      "",
			strLen:   5,
			expected: "00000",
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			res, err := Pad(test.str, test.strLen)
			if err != nil {
				t.Fatalf("unexpected non nil error: %v", err)
			}
			if res != test.expected {
				t.Errorf("result %s not equal to expected %s", res, test.expected)
			}
		})
	}
}

func TestPadWithError(t *testing.T) {
	type testCase struct {
		str    string
		strLen int
	}

	tests := map[string]testCase{
		"strLen less than length of input": {
			str:    "heLLo",
			strLen: 2,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			_, err := Pad(test.str, test.strLen)
			if err == nil {
				t.Fatal("expected non nil error")
			}
		})
	}
}

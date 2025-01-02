package challenge2

import "testing"

func TestConstraintValidation(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		out      string
		partial  bool
		expected bool
	}{
		{"Invalid length", "LL=", "3", false, false},
		{"Invalid L decode", "L", "23", false, false},
		{"Invalid R decode", "R", "32", false, false},
		{"Invalid = decode", "=", "32", false, false},
		{"Invalid input character", "Z", "", false, false},
		{"Invalid output character", "LL", "LL", false, false},
		{"Allow partial check", "LLRR=", "210", true, true},
		{"Valid input and output", "LLRR=", "210122", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValid(tt.in, tt.out, tt.partial)
			if result != tt.expected {
				t.Errorf("isValid(%q, %v, %t) = %t; want %t", tt.in, tt.out, tt.partial, result, tt.expected)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		expected string
	}{
		{"Decode case 1", "LLRR=", "210122"},
		{"Decode case 2", "==RLL", "000210"},
		{"Decode case 3", "=LLRR", "221012"},
		{"Decode case 4", "RRL=R", "012001"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := Decode(tt.in)
			if result.answer != tt.expected {
				t.Errorf("decode(%q) = %q; want %q", tt.in, result.answer, tt.expected)
			}
		})
	}
}

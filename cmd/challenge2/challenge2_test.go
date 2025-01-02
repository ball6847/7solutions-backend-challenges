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
		name       string
		in         string
		expected   string
		shouldFail bool
	}{
		{"Decode case 1", "LLRR=", "210122", false},
		{"Decode case 2", "==RLL", "000210", false},
		{"Decode case 3", "=LLRR", "221012", false},
		{"Decode case 4", "RRL=R", "012001", false},
		{"Empty input", "", "", true},
		{"Invalid input", "ZZZ", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Decode(tt.in)
			if tt.shouldFail {
				if err == nil {
					t.Errorf("Decode(%q) should have failed but did not", tt.in)
				}
			} else {
				if err != nil {
					t.Errorf("Decode(%q) failed with error: %v", tt.in, err)
				} else if result.optimal != tt.expected {
					t.Errorf("Decode(%q) = %q; want %q", tt.in, result.optimal, tt.expected)
				}
			}
		})
	}
}

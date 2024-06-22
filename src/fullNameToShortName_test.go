package main

import (
	"testing"
)

func TestFullNameToShortName(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		shouldError bool
	}{
		{"Magomadov Magomed Magomadovich", "Magomadov M. M.", false},
		{"John Doe Smith", "John D. S.", false},
		{"SingleWord", "", true},
		{"John Doe", "", true},
	}

	for _, test := range tests {
		output, err := fullNameToShortName(test.input)
		if test.shouldError {
			if err == nil {
				t.Errorf("fullNameToShortName(%q) expected error but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("fullNameToShortName(%q) unexpected error: %v", test.input, err)
			}
			if output != test.expected {
				t.Errorf("fullNameToShortName(%q) = %q; want %q", test.input, output, test.expected)
			}
		}
	}
}

package main

import (
	"github.com/Salladin95/doc-prompt/cmd"
	"testing"
)

func TestRetrieveFirstLetter(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Magomed", "M"},
		{"", ""},
		{"A", "A"},
	}

	for _, test := range tests {
		output := cmd.RetrieveFirstLetter(test.input)
		if output != test.expected {
			t.Errorf("RetrieveFirstLetter(%q) = %q; want %q", test.input, output, test.expected)
		}
	}
}

func TestRetrieveFirstWord(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Magomed Magomadov", "Magomed"},
		{"", ""},
		{"SingleWord", "SingleWord"},
	}

	for _, test := range tests {
		output := cmd.RetrieveFirstWord(test.input)
		if output != test.expected {
			t.Errorf("RetrieveFirstWord(%q) = %q; want %q", test.input, output, test.expected)
		}
	}
}

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

func TestFormatTimeOrDefault(t *testing.T) {
	tests := []struct {
		timeInput   string
		defaultTime string
		expected    string
	}{
		{"10 40", "17 часов 30 минут", "10 часов 40 минут"},
		{"", "17 часов 30 минут", "17 часов 30 минут"},
	}

	for _, test := range tests {
		output := cmd.FormatTimeOrDefault(test.timeInput, test.defaultTime)
		if output != test.expected {
			t.Errorf("FormatTimeOrDefault(%q, %q) = %q; want %q", test.timeInput, test.defaultTime, output, test.expected)
		}
	}
}

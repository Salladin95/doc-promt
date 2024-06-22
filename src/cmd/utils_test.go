package cmd

import "testing"

func TestIsOrdinanceSpecificKey(t *testing.T) {
	tests := []struct {
		input    UserInputKey
		expected bool
	}{
		{Decision, true},
		{NumberOfOrdinance, true},
		{DateOfEnactment, true},
		{FullName, false},
	}

	for _, test := range tests {
		output := IsOrdinanceSpecificKey(test.input)
		if output != test.expected {
			t.Errorf("IsOrdinanceSpecificKey(%q) = %v; want %v", test.input, output, test.expected)
		}
	}
}

func TestIsProtocolSpecificKey(t *testing.T) {
	tests := []struct {
		input    UserInputKey
		expected bool
	}{
		{NumberOfProtocol, false},
		{DateOfProtocol, false},
		{FullName, false},
	}

	for _, test := range tests {
		output := IsProtocolSpecificKey(test.input)
		if output != test.expected {
			t.Errorf("IsProtocolSpecificKey(%q) = %v; want %v", test.input, output, test.expected)
		}
	}
}

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
		output := RetrieveFirstLetter(test.input)
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
		output := RetrieveFirstWord(test.input)
		if output != test.expected {
			t.Errorf("RetrieveFirstWord(%q) = %q; want %q", test.input, output, test.expected)
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
		output := FormatTimeOrDefault(test.timeInput, test.defaultTime)
		if output != test.expected {
			t.Errorf("FormatTimeOrDefault(%q, %q) = %q; want %q", test.timeInput, test.defaultTime, output, test.expected)
		}
	}
}

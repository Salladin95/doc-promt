package main

import (
	"testing"
	"time"
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
		output := formatTimeOrDefault(test.timeInput, test.defaultTime)
		if output != test.expected {
			t.Errorf("formatTimeOrDefault(%q, %q) = %q; want %q", test.timeInput, test.defaultTime, output, test.expected)
		}
	}
}

func TestProcessDefaults(t *testing.T) {
	userInputs := map[UserInputKey]string{
		DateOfProtocol:  "15.06.2024г.",
		OfficialAddress: "123 Main St",
	}

	processDefaults(userInputs)

	dateOfProtocol, _ := time.Parse(dateFormat, "15.06.2024г.")
	expectedDateOfOrdinance := dateOfProtocol.AddDate(0, 0, 1).Format(dateFormat)
	expectedDateOfEnactment := dateOfProtocol.AddDate(0, 0, 11).Format(dateFormat)

	if userInputs[ActualAddress] != "123 Main St" {
		t.Errorf("Expected ActualAddress to be '123 Main St', got %q", userInputs[ActualAddress])
	}

	if userInputs[Occupation] != DefaultOccupation {
		t.Errorf("Expected Occupation to be '%s', got %q", DefaultOccupation, userInputs[Occupation])
	}

	if userInputs[DateOfOrdinance] != expectedDateOfOrdinance {
		t.Errorf("Expected DateOfOrdinance to be '%s', got %q", expectedDateOfOrdinance, userInputs[DateOfOrdinance])
	}

	if userInputs[DateOfEnactment] != expectedDateOfEnactment {
		t.Errorf("Expected DateOfEnactment to be '%s', got %q", expectedDateOfEnactment, userInputs[DateOfEnactment])
	}
}

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
		{NumberOfProtocol, true},
		{DateOfProtocol, true},
		{FullName, false},
	}

	for _, test := range tests {
		output := IsProtocolSpecificKey(test.input)
		if output != test.expected {
			t.Errorf("IsProtocolSpecificKey(%q) = %v; want %v", test.input, output, test.expected)
		}
	}
}

package cmd

import (
	"testing"
	"time"
)

func TestProcessDefaults(t *testing.T) {
	userInputs := map[UserInputKey]string{
		DateOfProtocol:  "15.06.2024г.",
		OfficialAddress: "123 Main St",
	}

	ProcessDefaults(userInputs)

	dateOfProtocol, _ := time.Parse(DateFormat, "15.06.2024г.")
	expectedDateOfOrdinance := dateOfProtocol.AddDate(0, 0, 1).Format(DateFormat)
	expectedDateOfEnactment := dateOfProtocol.AddDate(0, 0, 11).Format(DateFormat)

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

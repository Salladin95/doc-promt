package cmd

import (
	"log"
	"time"
)

func ProcessDefaults(userInputs map[UserInputKey]string) {
	dateOfProtocol, err := time.Parse(DateFormat, userInputs[DateOfProtocol])
	if err != nil {
		log.Panicln(err)
	}

	if userInputs[ActualAddress] == "" {
		userInputs[ActualAddress] = userInputs[OfficialAddress]
	}

	if userInputs[Occupation] == "" {
		userInputs[Occupation] = DefaultOccupation
	}

	if userInputs[DateOfOrdinance] == "" {
		userInputs[DateOfOrdinance] = dateOfProtocol.AddDate(0, 0, 1).Format(DateFormat)
	}

	userInputs[TimeOfOrdinance] = FormatTimeOrDefault(userInputs[TimeOfOrdinance], DefaultTimeOrdinance)
	userInputs[DateOfAccident] = dateOfProtocol.Format(DateFormat)
	userInputs[TimeOfAccident] = FormatTimeOrDefault(userInputs[TimeOfAccident], DefaultTimeAccident)

	if userInputs[Decision] == "" {
		userInputs[Decision] = DefaultDecision
	}

	dateOfOrdinance, err := time.Parse(DateFormat, userInputs[DateOfOrdinance])
	if err != nil {
		log.Panicln(err)
	}

	userInputs[DateOfEnactment] = dateOfOrdinance.AddDate(0, 0, 10).Format(DateFormat)
}
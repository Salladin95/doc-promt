package cmd

import (
	"fmt"
	"strings"
)

func IsOrdinanceSpecificKey(key UserInputKey) bool {
	switch key {
	case Decision, NumberOfOrdinance, DateOfEnactment:
		return true
	default:
		return false
	}
}

func IsProtocolSpecificKey(key UserInputKey) bool {
	switch key {
	case NumberOfProtocol, DateOfProtocol:
		return true
	default:
		return false
	}
}

func RetrieveFirstLetter(input string) string {
	if len(input) > 0 {
		return string([]rune(input)[0])
	}
	return ""
}

func RetrieveFirstWord(input string) string {
	if len(input) > 0 {
		fullNameParts := strings.Split(input, " ")
		return fullNameParts[0]
	}
	return ""
}

func FormatTimeOrDefault(timeInput, defaultTime string) string {
	if timeInput == "" {
		return defaultTime
	}
	timeParts := strings.Split(timeInput, " ")
	return fmt.Sprintf("%s часов %s минут", timeParts[0], timeParts[1])
}

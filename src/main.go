package main

import (
	"bufio"
	"fmt"
	"github.com/lukasjarosch/go-docx"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	userInputs := make(map[UserInputKey]string)

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	for _, pair := range placeholders {
		key := pair[0].(UserInputKey)
		prompt := pair[1].(string)
		fmt.Println(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if key == FullName {
			userInputs[ShortName] = fullNameToShortName(input)
		}
		userInputs[key] = input
	}

	protocolReplaceMap := make(docx.PlaceholderMap)
	ordinanceReplaceMap := make(docx.PlaceholderMap)

	for key, value := range userInputs {
		formattedKey := fmt.Sprintf("{%s}", key)
		switch true {
		case IsProtocolSpecificKey(key):
			protocolReplaceMap[formattedKey] = value
		case IsOrdinanceSpecificKey(key):
			ordinanceReplaceMap[formattedKey] = value
		default:
			protocolReplaceMap[formattedKey] = value
			ordinanceReplaceMap[formattedKey] = value
		}
	}

	// Construct the absolute path to the template
	protocolTemplatePath := filepath.Join(cwd, "templates", "protocol_template.docx")
	ordinanceTemplatePath := filepath.Join(cwd, "templates", "ordinance_template.docx")

	// Load the template document
	protocol, err := docx.Open(protocolTemplatePath)
	if err != nil {
		fmt.Println("Error opening  protocol:", err)
		return
	}

	// Load the template document
	ordinance, err := docx.Open(ordinanceTemplatePath)
	if err != nil {
		fmt.Println("Error opening ordinance:", err)
		return
	}

	if err := protocol.ReplaceAll(protocolReplaceMap); err != nil {
		panic(err)
	}

	if err := ordinance.ReplaceAll(ordinanceReplaceMap); err != nil {
		panic(err)
	}

	folderPath := filepath.Join(
		cwd,
		"filled_documents",
		fmt.Sprintf("%s %s", RetrieveFirstWord(userInputs[FullName]), userInputs[DateOfAccident]),
	)

	// Create directory with read and write permissions for the user
	if err := os.MkdirAll(folderPath, 0755); err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	filledProtocolPath := filepath.Join(folderPath, "filled_protocol.docx")
	fmt.Println(filledProtocolPath)
	if err := protocol.WriteToFile(filledProtocolPath); err != nil {
		fmt.Println("Error saving protocol:", err)
		return
	}

	filledOrdinancePath := filepath.Join(folderPath, "filled_ordinance.docx")
	fmt.Println(filledProtocolPath)
	if err := protocol.WriteToFile(filledOrdinancePath); err != nil {
		fmt.Println("Error saving ordinance:", err)
		return
	}

	fmt.Println("Documents created successfully:", filledProtocolPath)
}

// UserInputKey represents the valid keys for the placeholders map
type UserInputKey string

const (
	FullName          UserInputKey = "fullName"
	ShortName         UserInputKey = "shortName"
	Birthday          UserInputKey = "birthday"
	PlaceOfBirth      UserInputKey = "placeOfBirth"
	OfficialAddress   UserInputKey = "officialAddress"
	ActualAddress     UserInputKey = "actualAddress"
	Identifier        UserInputKey = "identifier"
	Occupation        UserInputKey = "occupation"
	NumberOfProtocol  UserInputKey = "numberOfProtocol"
	DateOfProtocol    UserInputKey = "dateOfProtocol"
	DateOfAccident    UserInputKey = "dateOfAccident"
	TimeOfAccident    UserInputKey = "timeOfAccident"
	NumberOfOrdinance UserInputKey = "numberOfOrdinance"
	DateOfOrdinance   UserInputKey = "dateOfOrdinance"
	TimeOfOrdinance   UserInputKey = "timeOfOrdinance"
	Decision          UserInputKey = "decision"
	DateOfEnactment   UserInputKey = "dateOfEnactment"
)

// Example usage
var placeholders = [][]interface{}{
	{NumberOfProtocol, "Введите № протокола: "},
	{DateOfProtocol, "Введите дату регистрации протокола, в следующем формате - 30.12.2024.: "},
	{FullName, "Введите полное имя - Магомадов Магомед Магомедович: "},
	{Birthday, "Введите дату рождения, в следующем формате - 30.12.1954г.: "},
	{PlaceOfBirth, "Введите место рождения(как в паспорте): "},
	{OfficialAddress, "Введите адресс регистрации: "},
	{ActualAddress, "Введите фактический адресс: "},
	{Identifier, "Введите документ, удостоверяющий личность - 'Паспорт серия 9610 № 224309 выдан Отделом УФМС России по ЧР в Гудермесском районе: "},
	{DateOfAccident, "Введите дату происшествия, в следующем формате - 30.12.2024г: "},
	{TimeOfAccident, "Введите время происшествия, в следующем формате - 10 40: "},
	{Occupation, "Введите должность - \"Директор\" || \"Гражданин\": "},
	{NumberOfOrdinance, "Введите № постановления: "},
	{DateOfOrdinance, "Введите дату рассмотрения протокола(Дата регистрации постановления), в следующем формате 30.12.2024г: "},
	{TimeOfOrdinance, "Введите время рассмотрения протокла, в следующем формате - 10 40: "},
	{Decision, "Введите решение постановления в родительноме падеже, например - \"штрафа в размере 5000 (ПЯТЬ ТЫСЯЧ РУБЛЕЙ)\" || \"Предупреждения\": "},
	{DateOfEnactment, "Введите дату вступления постановления в законную силу(10 дней со дня вынесения) в следующем формате - 30.12.2024г."},
}

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

func fullNameToShortName(fullName string) string {
	fullNameParts := strings.Split(fullName, " ")
	return fmt.Sprintf(
		"%s %s. %s.",
		fullNameParts[0],
		RetrieveFirstLetter(fullNameParts[1]),
		RetrieveFirstLetter(fullNameParts[2]),
	)
}

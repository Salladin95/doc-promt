package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/lukasjarosch/go-docx"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const dateFormat = "02.01.2006г."

// Default values
const (
	DefaultOccupation    = "Гражданин"
	DefaultTimeOrdinance = "17 часов 30 минут"
	DefaultTimeAccident  = "11 часов 30 минут"
	DefaultDecision      = "Предупреждения"
)

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

var Placeholders = [][]interface{}{
	{NumberOfProtocol, "Введите № протокола: "},
	{DateOfProtocol, fmt.Sprintf("Введите дату регистрации протокола, в следующем формате - %s: ", dateFormat)},
	{FullName, "Введите полное имя - Магомадов Магомед Магомедович: "},
	{Birthday, "Введите дату рождения, в следующем формате - 30.12.1954г.: "},
	{PlaceOfBirth, "Введите место рождения(как в паспорте): "},
	{OfficialAddress, "Введите адрес регистрации: "},
	{ActualAddress, "Введите фактический адрес (по умолчанию - адрес регистрации): "},
	{Identifier, "Введите документ, удостоверяющий личность - 'Паспорт серия 9610 № 224309 выдан Отделом УФМС России по ЧР в Гудермесском районе 20.12.2012г.: "},
	{DateOfAccident, fmt.Sprintf("Введите дату происшествия, в следующем формате - %s (по умолчанию - дата регистрации протокола): ", dateFormat)},
	{TimeOfAccident, "Введите время происшествия, в следующем формате - 10 40 (по умолчанию - 11 30): "},
	{Occupation, "Введите должность - \"Директор\" || \"Гражданин\" (по умолчанию Гражданин): "},
	{NumberOfOrdinance, "Введите № постановления: "},
	{DateOfOrdinance, fmt.Sprintf("Введите дату рассмотрения протокола (Дата регистрации постановления), в следующем формате - %s  (по умолчанию следующий день от даты регистрации протокола): ", dateFormat)},
	{TimeOfOrdinance, "Введите время рассмотрения протокола, в следующем формате - 10 40  (по умолчанию - 17 30): "},
	{Decision, "Введите решение постановления в родительноме падеже, например - \"штрафа в размере 5000 (ПЯТЬ ТЫСЯЧ РУБЛЕЙ)\" || \"Предупреждения\"  (по умолчанию - Предупреждения): "},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	userInputs := gatherUserInputs(reader)
	processDefaults(userInputs)

	err := generateDocuments(userInputs)
	if err != nil {
		log.Fatalf("Error generating documents: %v", err)
	}
	fmt.Println("Documents created successfully")
}

func gatherUserInputs(reader *bufio.Reader) map[UserInputKey]string {
	userInputs := make(map[UserInputKey]string)

	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println("!!!!!!!!! Если у поля есть значение по умолчанию, оно будет использовано в случае пропуска заполнения данного поля !!!!!!!!!")
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println()

	for _, pair := range Placeholders {
		key := pair[0].(UserInputKey)
		prompt := pair[1].(string)
		fmt.Println(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		userInputs[key] = input
	}

	shortName, err := fullNameToShortName(userInputs[FullName])
	if err != nil {
		log.Panicln(err)
	}
	userInputs[ShortName] = shortName
	return userInputs
}

func processDefaults(userInputs map[UserInputKey]string) {
	dateOfProtocol, err := time.Parse(dateFormat, userInputs[DateOfProtocol])
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
		userInputs[DateOfOrdinance] = dateOfProtocol.AddDate(0, 0, 1).Format(dateFormat)
	}

	userInputs[TimeOfOrdinance] = formatTimeOrDefault(userInputs[TimeOfOrdinance], DefaultTimeOrdinance)
	userInputs[DateOfAccident] = dateOfProtocol.Format(dateFormat)
	userInputs[TimeOfAccident] = formatTimeOrDefault(userInputs[TimeOfAccident], DefaultTimeAccident)

	if userInputs[Decision] == "" {
		userInputs[Decision] = DefaultDecision
	}

	dateOfOrdinance, err := time.Parse(dateFormat, userInputs[DateOfOrdinance])
	if err != nil {
		log.Panicln(err)
	}

	userInputs[DateOfEnactment] = dateOfOrdinance.AddDate(0, 0, 10).Format(dateFormat)
}

func formatTimeOrDefault(timeInput, defaultTime string) string {
	if timeInput == "" {
		return defaultTime
	}
	timeParts := strings.Split(timeInput, " ")
	return fmt.Sprintf("%s часов %s минут", timeParts[0], timeParts[1])
}

func generateDocuments(userInputs map[UserInputKey]string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting current working directory: %v", err)
	}

	protocolReplaceMap, ordinanceReplaceMap := createReplacementMaps(userInputs)
	folderPath := createFolderPath(cwd, userInputs)

	err = createFilledDocument(filepath.Join(cwd, "templates", "protocol_template.docx"), protocolReplaceMap, filepath.Join(folderPath, "filled_protocol.docx"))
	if err != nil {
		return err
	}

	err = createFilledDocument(filepath.Join(cwd, "templates", "ordinance_template.docx"), ordinanceReplaceMap, filepath.Join(folderPath, "filled_ordinance.docx"))
	if err != nil {
		return err
	}

	return nil
}

func createReplacementMaps(userInputs map[UserInputKey]string) (docx.PlaceholderMap, docx.PlaceholderMap) {
	protocolReplaceMap := make(docx.PlaceholderMap)
	ordinanceReplaceMap := make(docx.PlaceholderMap)

	for key, value := range userInputs {
		formattedKey := fmt.Sprintf("{%s}", key)
		switch {
		case IsProtocolSpecificKey(key):
			protocolReplaceMap[formattedKey] = value
		case IsOrdinanceSpecificKey(key):
			ordinanceReplaceMap[formattedKey] = value
		default:
			protocolReplaceMap[formattedKey] = value
			ordinanceReplaceMap[formattedKey] = value
		}
	}

	return protocolReplaceMap, ordinanceReplaceMap
}

func createFilledDocument(templatePath string, replaceMap docx.PlaceholderMap, outputPath string) error {
	doc, err := docx.Open(templatePath)
	if err != nil {
		return fmt.Errorf("opening template %s: %v", templatePath, err)
	}

	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		return fmt.Errorf("replacing Placeholders: %v", err)
	}

	err = doc.WriteToFile(outputPath)
	if err != nil {
		return fmt.Errorf("saving document: %v", err)
	}

	return nil
}

func createFolderPath(cwd string, userInputs map[UserInputKey]string) string {
	folderPath := filepath.Join(
		cwd,
		"filled_documents",
		fmt.Sprintf("%s %s", RetrieveFirstWord(userInputs[FullName]), userInputs[DateOfAccident]),
	)

	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	return folderPath
}

func fullNameToShortName(fullName string) (string, error) {
	fullNameParts := strings.Split(fullName, " ")
	if len(fullNameParts) < 3 {
		return "", errors.New("full name too short")
	}
	return fmt.Sprintf(
		"%s %s. %s.",
		fullNameParts[0],
		RetrieveFirstLetter(fullNameParts[1]),
		RetrieveFirstLetter(fullNameParts[2]),
	), nil
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

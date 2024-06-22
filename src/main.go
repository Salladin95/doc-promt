package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Salladin95/doc-prompt/cmd"
	"github.com/lukasjarosch/go-docx"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	userInputs := gatherUserInputs(reader)
	cmd.ProcessDefaults(userInputs)

	if err := generateDocuments(userInputs); err != nil {
		log.Fatalf("Error generating documents: %v", err)
	}
	fmt.Println("Documents created successfully")
}

func gatherUserInputs(reader *bufio.Reader) map[cmd.UserInputKey]string {
	userInputs := make(map[cmd.UserInputKey]string)

	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println("!!!!!!!!! Если у поля есть значение по умолчанию, оно будет использовано в случае пропуска заполнения данного поля !!!!!!!!!")
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	fmt.Println()

	for _, pair := range cmd.Placeholders {
		key := pair[0].(cmd.UserInputKey)
		prompt := pair[1].(string)
		fmt.Println(prompt)
		userInput, _ := reader.ReadString('\n')
		userInputs[key] = strings.TrimSpace(userInput)
	}

	shortName, err := fullNameToShortName(userInputs[cmd.FullName])
	if err != nil {
		log.Panicln(err)
	}
	userInputs[cmd.ShortName] = shortName
	return userInputs
}

func generateDocuments(userInputs map[cmd.UserInputKey]string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting current working directory: %v", err)
	}

	protocolReplaceMap, ordinanceReplaceMap := createReplacementMaps(userInputs)
	folderPath := createFolderPath(cwd, generateFolderPathSuffix(userInputs))

	if err := createFilledDocument(filepath.Join(cwd, "templates", "protocol_template.docx"), protocolReplaceMap, filepath.Join(folderPath, fmt.Sprintf("Проткол № %s.docx", userInputs[cmd.NumberOfProtocol]))); err != nil {
		return err
	}

	if err := createFilledDocument(filepath.Join(cwd, "templates", "ordinance_template.docx"), ordinanceReplaceMap, filepath.Join(folderPath, fmt.Sprintf("Постановление № %s.docx", userInputs[cmd.NumberOfOrdinance]))); err != nil {
		return err
	}

	return nil
}

func createReplacementMaps(userInputs map[cmd.UserInputKey]string) (docx.PlaceholderMap, docx.PlaceholderMap) {
	protocolReplaceMap := make(docx.PlaceholderMap)
	ordinanceReplaceMap := make(docx.PlaceholderMap)

	for key, value := range userInputs {
		formattedKey := fmt.Sprintf("{%s}", key)
		switch {
		case cmd.IsProtocolSpecificKey(key):
			protocolReplaceMap[formattedKey] = value
		case cmd.IsOrdinanceSpecificKey(key):
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

func createFolderPath(cwd, suffix string) string {
	folderPath := filepath.Join(
		cwd,
		"materials",
		suffix,
	)

	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	return folderPath
}

func generateFolderPathSuffix(userInputs cmd.UserInputsMap) string {
	return fmt.Sprintf(
		"%s %s",
		cmd.RetrieveFirstWord(userInputs[cmd.FullName]),
		userInputs[cmd.DateOfAccident],
	)
}

func fullNameToShortName(fullName string) (string, error) {
	fullNameParts := strings.Split(fullName, " ")
	if len(fullNameParts) < 3 {
		return "", errors.New("full name too short")
	}
	return fmt.Sprintf(
		"%s %s. %s.",
		fullNameParts[0],
		cmd.RetrieveFirstLetter(fullNameParts[1]),
		cmd.RetrieveFirstLetter(fullNameParts[2]),
	), nil
}

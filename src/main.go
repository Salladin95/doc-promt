package main

import (
	"bufio"
	"fmt"
	"github.com/lukasjarosch/go-docx"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for user input
	fmt.Print("Enter the number: ")
	number, _ := reader.ReadString('\n')
	number = strings.TrimSpace(number)

	fmt.Print("Enter the date: ")
	date, _ := reader.ReadString('\n')
	date = strings.TrimSpace(date)

	fmt.Print("Enter full name: ")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName)

	//fullNameParts := strings.Split(fullName, " ")
	//
	//shortName := fmt.Sprintf(
	//	"%s %s. %s.",
	//	fullNameParts[0],
	//	string(fullNameParts[1][0]),
	//	string(fullNameParts[2][0]),
	//)

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// Construct the absolute path to the template
	protocolTemplatePath := filepath.Join(cwd, "..", "protocol_template.docx")

	// replaceMap is a key-value map whereas the keys
	// represent the placeholders without the delimiters
	replaceMap := docx.PlaceholderMap{
		"{number}":    "12",
		"{date}":      "12/12/2024",
		"{fullName}":  "Shamurzaev Khalid Abubakarovich",
		"{shortName}": "Shamurzaev Kh. A.",
	}

	// Load the template document
	doc, err := docx.Open(protocolTemplatePath)
	if err != nil {
		fmt.Println("Error opening document:", err)
		return
	}

	if err := doc.ReplaceAll(replaceMap); err != nil {
		fmt.Println(err)
	}

	filledProtocolPath := filepath.Join(cwd, "..", "filled_protocol.docx")
	err = doc.WriteToFile(filledProtocolPath)
	if err != nil {
		fmt.Println("Error saving document:", err)
		return
	}

	fmt.Println("Document created successfully:", filledProtocolPath)
}

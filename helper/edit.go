package helper

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func EditFile() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	projectDir := filepath.Join(userHomeDir, "tasks")

	var target string
	fmt.Print("Enter the file name to edit: ")
	fmt.Scanln(&target)

	filePath := filepath.Join(projectDir, target)

	file, openErr := os.OpenFile(filePath, os.O_RDWR, 0o644)
	if openErr != nil {
		log.Fatal(openErr.Error())
	}
	defer file.Close()

	var content string

	fmt.Print("Append content: ")
	fmt.Scanln(&content)

	_, writeErr := file.WriteString(content + "\n")
	if writeErr != nil {
		log.Fatal(writeErr.Error())
	}

	fmt.Println("Content appended successfully.")
	os.Exit(1)
}

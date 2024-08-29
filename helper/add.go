package helper

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func AddNewTask() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to read user home directory")
		log.Fatal(err.Error())
	}

	var fileName string
	var content string

	fmt.Print("Enter file name with extentions: ")
	fmt.Scanln(&fileName)

	fmt.Print("add content to this file: ")
	fmt.Scanln(&content)

	target := filepath.Join(userHomeDir, "tasks")

	location := filepath.Join(target, fileName)

	writeErr := os.WriteFile(location, []byte(content+"\n"), 0o644)
	if writeErr != nil {
		log.Fatal(writeErr.Error())
	}
	fmt.Println("New task has been added to ", target)
}

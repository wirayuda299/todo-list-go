package helper

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ListAllTask() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
	}
	location := filepath.Join(userHomeDir, "tasks")

	dirEntries, error := os.ReadDir(location)
	if error != nil {
		log.Fatal(error.Error())
	}

	for _, de := range dirEntries {
		fmt.Println(de.Name())
	}
}

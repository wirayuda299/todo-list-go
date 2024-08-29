package helper

import (
	"fmt"
	"os"
)

func InitProject() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())
	}

	errors := os.Mkdir(userHomeDir+"/tasks", 0o777)

	if errors != nil && os.IsExist(errors) {
		fmt.Println("Directory exists")
		return
	} else if errors != nil && !os.IsExist(errors) {
		fmt.Println(errors.Error())
		return
	} else {
		fmt.Println("Success, add your first task now!")
	}
}

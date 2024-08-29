/*
Copyright © 2024 wirayuda <wirayuda233@gmail.com>
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Task management using CLI",
	Long: `Task management tool using Command Line Interface`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

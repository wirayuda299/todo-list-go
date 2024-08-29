/*
Copyright © 2024 wirayuda <wirayuda233@gmail.com>
*/


package cmd

import (
	"tasks/helper"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Long:  `Add new task `,
	Run: func(cmd *cobra.Command, args []string) {
		helper.AddNewTask()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

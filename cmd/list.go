/*
Copyright © 2024 wirayuda <wirayuda233@gmail.com>
*/
package cmd

import (
	"tasks/helper"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all task",
	Long:  `Show all task list`,

	Run: func(cmd *cobra.Command, args []string) {
		helper.ListAllTask()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

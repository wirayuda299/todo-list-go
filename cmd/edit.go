/*
Copyright © 2024 wirayuda <wirayuda233@gmail.com>
*/

package cmd

import (
	"tasks/helper"

	"github.com/spf13/cobra"
)


var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a file",
	Run: func(cmd *cobra.Command, args []string) {
    helper.EditFile()
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}

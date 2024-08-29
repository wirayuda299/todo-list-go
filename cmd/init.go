/*
Copyright © 2024 wirayuda <wirayuda233@gmail.com>
*/

package cmd

import (
	"tasks/helper"

	"github.com/spf13/cobra"
)


var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init Project",
	Long:  `Init project so you can add task later`,
	Run: func(cmd *cobra.Command, args []string) {
    helper.InitProject()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

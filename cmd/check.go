package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if latest AS info conflict with previous",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			return
		}
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			cmd.Help()
			return
		}
	},
}

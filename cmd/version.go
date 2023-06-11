package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"registry-father/config"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.Version)
	},
}

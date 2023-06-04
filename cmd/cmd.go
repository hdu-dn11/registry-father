package cmd

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"os"
	"registry-father/pages"
)

var rootCmd = &cobra.Command{
	Use:   "dn11",
	Short: "dn11 is a tool to help you manage AS in DN11",
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()
		if err := app.SetRoot(pages.Load(app), true).Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

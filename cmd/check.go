package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"registry-father/services"
	"sort"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if latest AS info conflict with previous",
	Run: func(cmd *cobra.Command, args []string) {
		list, err := services.GetASInfoList()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Check if latest AS info conflict with previous
		if len(list) <= 1 {
			return
		}
		sort.SliceStable(list, func(i, j int) bool {
			return list[i].UpdatedAt.After(list[j].UpdatedAt)
		})
		for i := 1; i < len(list); i++ {
			if services.CheckCIDRConflict(list[0], list[i]) {
				fmt.Printf("AS%d conflict with AS%d\n", list[0].ASN, list[i].ASN)
				os.Exit(1)
			}
		}
	},
}

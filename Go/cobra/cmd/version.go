package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:  "version",
	Long: "FIGHTING",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ryan v1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

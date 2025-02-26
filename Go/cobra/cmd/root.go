package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rayn",
	Short: "Rayn is a fking cool guy m3",
	Long:  `fighting`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run ryan...")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)

	}
}

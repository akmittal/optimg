package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "optimg",
	Short: "Optimize Images for serving smaller images",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello world")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

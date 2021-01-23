package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows version of command",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("optimg version: 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(VersionCmd)
}

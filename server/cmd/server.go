package cmd

import (
	"log"

	"github.com/akmittal/optimg/server/pkg/application"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Long:  `Start server`,
	Run: func(cmd *cobra.Command, args []string) {
		app, err := application.Get()
		if err != nil {
			log.Fatal(err.Error())
		}
		err = app.Start()
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringP("host", "a", "localhost", "Application host where to run")
	serverCmd.Flags().StringP("port", "p", "8000", "Optimg dashboard port")

}

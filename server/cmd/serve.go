package cmd

import (
	"fmt"
	"log"

	"github.com/akmittal/optimg/server/pkg/application"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start static server",
	Long:  `Start static server`,
	Run: func(cmd *cobra.Command, args []string) {
		app, err := application.Get()
		if err != nil {
			log.Fatal(err.Error())
		}
		directory, _ := cmd.Flags().GetString("directory")
		fmt.Printf("App running on %v", app.Cfg.GetAppHost())
		err = app.StartStaticServer(directory)
		if err != nil {
			log.Fatal(err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("host", "a", "localhost", "Application host where to run")
	serveCmd.Flags().StringP("port", "p", "8000", "Optimg dashboard port")
	serveCmd.Flags().StringP("directory", "d", "/tmp", "Directory which is to be served")
	serveCmd.MarkFlagRequired("directory")

}

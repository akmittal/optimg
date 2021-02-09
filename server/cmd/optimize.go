package cmd

import (
	"fmt"
	"log"

	"github.com/akmittal/optimg/server/pkg/application"
	"github.com/akmittal/optimg/server/pkg/operation"
	"github.com/spf13/cobra"
)

var optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Optimize Images for serving smaller images",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		app, err := application.Get()
		sourcePath, _ := cmd.Flags().GetString("sourcePath")
		targetPath, _ := cmd.Flags().GetString("targetPath")
		copyUnknown, _ := cmd.Flags().GetBool("copyUnknown")
		formats, _ := cmd.Flags().GetStringSlice("format")
		qualities, _ := cmd.Flags().GetIntSlice("quality")
		if len(formats) != len(qualities) {
			log.Fatal("Number of formats and qualities should be equal")
			return
		}
		var transformations []operation.Transformation
		for idx, format := range formats {
			transformation := operation.Transformation{Format: operation.FormatMapping[format], Quality: qualities[idx]}
			transformations = append(transformations, transformation)
		}
		fmt.Println(sourcePath, targetPath, copyUnknown, formats, qualities, transformations)
		opr, err := operation.Get(sourcePath, targetPath, copyUnknown, false, transformations)
		if err != nil {
			log.Fatalln("Error creating operation", err.Error())
		}
		err = opr.Process(app.DB.Client)
		if err != nil {
			log.Fatalln("Error creating operation", err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(optimizeCmd)
	optimizeCmd.Flags().StringP("sourcePath", "s", "/tmp", "Source Image Path")
	optimizeCmd.Flags().StringP("targetPath", "t", "/tmp", "Image Path where to store images")
	// optimizeCmd.Flags().BoolP("copyUnknown", "u", false, "Copy no-image/unsuppported files")
	optimizeCmd.Flags().StringSliceP("format", "f", []string{}, "Image format")
	optimizeCmd.Flags().IntSliceP("quality", "q", []int{}, "Image quality")
	optimizeCmd.MarkFlagRequired("sourcePath")
	optimizeCmd.MarkFlagRequired("targetPath")
	optimizeCmd.MarkFlagRequired("format")
	optimizeCmd.MarkFlagRequired("quality")
}

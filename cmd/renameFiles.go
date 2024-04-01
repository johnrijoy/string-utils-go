/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"example.com/user/string-utils/app"
	"example.com/user/string-utils/utils"
	"github.com/spf13/cobra"
)

// renameFilesCmd represents the renameFiles command
var renameFilesCmd = &cobra.Command{
	Use:   "renameFiles [basePath] regexFile newfileName",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		utils.DebugLn("renameFiles called")

		var basePath, filePattern, newFileName string

		if len(args) > 2 {
			basePath = args[0]
			filePattern = args[1]
			newFileName = args[2]
		} else {
			var err error
			if basePath, err = os.Getwd(); err != nil {
				utils.HandleExit(err)
			}
			filePattern = args[0]
			newFileName = args[1]
		}

		utils.InfoLn("Base path:", basePath)
		utils.InfoLn("File pattern:", filePattern)
		utils.InfoLn("File name template:", newFileName)

		if err := app.RenameFiles(basePath, filePattern, newFileName); err != nil {
			utils.HandleExit(err)
		}

		utils.DebugLn("renameFiles End")
	},
}

func init() {
	rootCmd.AddCommand(renameFilesCmd)

	// Here you will define your flags and configuration settings.

}

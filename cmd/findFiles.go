/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"example.com/user/string-utils/utils"
	"github.com/spf13/cobra"
)

// findFilesCmd represents the findFiles command
var findFilesCmd = &cobra.Command{
	Use:   "findFiles filePath",
	Short: "Command to check all files that match the pattern",
	Long:  `Command to check all files that match the pattern`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.DebugLn("findFiles called")

		filePath := args[0]

		if isInputPathFileMode {
			utils.InfoLn("File Path Mode: Input Path File")
			filePathList = utils.ReadInputPathFile(filePath)
		} else {
			utils.InfoLn("File Path Mode: Single")
			filePathList = append(filePathList, filePath)
		}

		utils.InfoLn("File Pattern: ", filePathList)

		resultList := utils.FindFilesFromGlobPatterns(filePathList)

		for _, result := range resultList {
			utils.InfoLn(result)
		}

		utils.DebugLn("findFiles End")
	},
}

func init() {
	rootCmd.AddCommand(findFilesCmd)

	// Flags and configuration settings
}

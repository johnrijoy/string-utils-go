/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"example.com/user/string-utils/app"
	"example.com/user/string-utils/utils"
	"github.com/spf13/cobra"
)

// FindLinesCmd represents the FindLines command
var findLinesCmd = &cobra.Command{
	Use:   "findLines",
	Short: "Find lines that matches",
	Long:  `Find all lines that match the given pattern in multiple files`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		utils.DebugLn("findLines called")

		regexExp := args[0]
		filePath := args[1]

		if isInputPathFileMode {
			utils.DebugLn("File Path Mode: Input Path File")
			filePathList = utils.ReadInputPathFile(filePath)
		} else {
			utils.DebugLn("File Path Mode: Single")
			filePathList = append(filePathList, filePath)
		}

		utils.InfoLn("Regex: ", regexExp)
		utils.InfoLn("Path: ", filePathList)

		// resultList := app.FindAllFromFile(regexExp, filePath)
		if !submatchFlag {
			resultList := app.FindAllLinesFromGlobPattern(regexExp, filePathList)
			utils.PrintOccurenceMap(resultList)
		} else {
			resultList := app.FindAllLinesSubmatchFromGlobPattern(regexExp, filePathList)
			utils.PrintSubmatchMap(resultList)
		}

		utils.DebugLn("findLines End")
	},
}

func init() {
	rootCmd.AddCommand(findLinesCmd)

	// Flags and configuration settings

	findLinesCmd.Flags().BoolVarP(&submatchFlag, "submatch", "s", false, "Enable capture groups")
}

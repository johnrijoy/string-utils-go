/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/johnrijoy/string-utils-go/app"
	"github.com/johnrijoy/string-utils-go/utils"
	"github.com/spf13/cobra"
)

// findAllCmd represents the findAll command
var findAllCmd = &cobra.Command{
	Use:   "findAll regex filePath",
	Short: "find all that matches",
	Long:  `Find all lines that match the given pattern in multiple files`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		utils.DebugLn("findAll called")

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
		findMatch := app.FindAllFromGlobPattern
		findSubMatch := app.FindAllSubmatchFromGlobPattern
		if lineNumberFlag {
			findMatch = app.FindAllLinesFromGlobPattern
			findSubMatch = app.FindAllLinesSubmatchFromGlobPattern
		}

		if !submatchFlag {
			resultList := findMatch(regexExp, filePathList)
			app.PrintOccurenceMap(resultList)
		} else {
			resultList := findSubMatch(regexExp, filePathList)
			app.PrintSubmatchMap(resultList)
		}

		utils.DebugLn("findAll End")
	},
}

func init() {
	rootCmd.AddCommand(findAllCmd)

	// Flags and configuration settings

	findAllCmd.Flags().BoolVarP(&submatchFlag, "submatch", "s", false, "Enable capture groups")
	findAllCmd.Flags().BoolVarP(&lineNumberFlag, "lineNumber", "l", false, "Enable line number")
}

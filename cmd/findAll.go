/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"example.com/user/string-utils/app"
	"example.com/user/string-utils/utils"
	"github.com/spf13/cobra"
)

// findAllCmd represents the findAll command
var findAllCmd = &cobra.Command{
	Use:   "findAll regex filePath",
	Short: "find all that matches",
	Long:  `Find all lines that match the given pattern in multiple files`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("findAll called")

		regexExp := args[0]
		filePath := args[1]

		if isInputPathFileMode {
			fmt.Println("File Path Mode: Input Path File")
			filePathList = utils.ReadInputPathFile(filePath)
		} else {
			fmt.Println("File Path Mode: Single")
			filePathList = append(filePathList, filePath)
		}

		fmt.Println("Regex: ", regexExp)
		fmt.Println("Path: ", filePathList)

		// resultList := app.FindAllFromFile(regexExp, filePath)
		if !submatchFlag {
			resultList := app.FindAllFromGlobPattern(regexExp, filePathList)
			utils.PrintOccurenceMap(resultList)
		} else {
			resultList := app.FindAllSubmatchFromGlobPattern(regexExp, filePathList)
			utils.PrintSubmatchMap(resultList)
		}

		fmt.Println("findAll End")
	},
}

func init() {
	rootCmd.AddCommand(findAllCmd)

	// Flags and configuration settings

	findAllCmd.Flags().BoolVarP(&submatchFlag, "submatch", "s", false, "Enable capture groups")
}

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

// replaceAllCmd represents the replaceAll command
var replaceAllCmd = &cobra.Command{
	Use:   "replaceAll REGEX FILE_PATH REPLACE_TEXT",
	Short: "replace all that matches",
	Long:  `Replace all lines that match the given pattern in multiple files`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("replaceAll called")
		regexExp := args[0]
		filePath := args[1]
		replaceText := args[2]

		fmt.Println("Regex: ", regexExp)
		fmt.Println("File: ", filePath)
		fmt.Println("Replace Text: ", replaceText)

		if !submatchFlag {
			resultList := app.ResplaceAllInGlobPattern(regexExp, filePath, replaceText)
			utils.PrintOccurenceMap(resultList)
		} else {
			resultList := app.ResplaceAllSubmatchesInGlobPattern(regexExp, filePath, replaceText)
			utils.PrintSubmatchMap(resultList)
		}

		fmt.Println("replaceAll End")
	},
}

func init() {
	rootCmd.AddCommand(replaceAllCmd)

	// Flags and configuration settings

	replaceAllCmd.Flags().BoolVarP(&submatchFlag, "submatch", "s", false, "Enable capture groups")
}

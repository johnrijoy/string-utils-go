/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/johnrijoy/string-utils-go/app"
	"github.com/johnrijoy/string-utils-go/utils"
	"github.com/spf13/cobra"
)

var isTemplateMode bool

// replaceAllCmd represents the replaceAll command
var replaceAllCmd = &cobra.Command{
	Use:   "replaceAll regex replaceText filePath",
	Short: "replace all that matches",
	Long:  `Replace all lines that match the given pattern in multiple files`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("replaceAll called")

		regexExp := args[0]
		replaceText := args[1]
		filePath := args[2]

		if isInputPathFileMode {
			utils.InfoLn("File Path Mode: Input Path File")
			filePathList = utils.ReadInputPathFile(filePath)
		} else {
			utils.InfoLn("File Path Mode: Single")
			filePathList = append(filePathList, filePath)
		}

		utils.InfoLn("Regex: ", regexExp)
		utils.InfoLn("File: ", filePathList)
		utils.InfoLn("Replace Text: ", replaceText)

		if !submatchFlag {
			resultList := app.ReplaceAllInGlobPatterns(regexExp, filePathList, replaceText, isTemplateMode)
			app.PrintOccurenceMap(resultList)
		} else {
			resultList := app.ReplaceAllSubmatchesInGlobPatterns(regexExp, filePathList, replaceText, isTemplateMode)
			app.PrintChangeMap(resultList)
		}

		utils.DebugLn("replaceAll End")
	},
}

func init() {
	rootCmd.AddCommand(replaceAllCmd)

	// Flags and configuration settings

	replaceAllCmd.Flags().BoolVarP(&submatchFlag, "submatch", "s", false, "Enable capture groups")
	replaceAllCmd.Flags().BoolVarP(&isTemplateMode, "template", "t", false, "Enable replace using template file")
}

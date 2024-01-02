/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"example.com/user/string-utils/utils"
	"github.com/spf13/cobra"
)

// findFilesCmd represents the findFiles command
var findFilesCmd = &cobra.Command{
	Use:   "findFiles",
	Short: "Command to check all files that match the pattern",
	Long:  `Command to check all files that match the pattern`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		filePath := args[0]

		fmt.Println("File Pattern: ", filePath)

		resultList := utils.FindFilesFromGlobPattern(filePath)

		for _, result := range resultList {
			fmt.Println(result)
		}

		fmt.Println("findFiles End")
	},
}

func init() {
	rootCmd.AddCommand(findFilesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findFilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findFilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

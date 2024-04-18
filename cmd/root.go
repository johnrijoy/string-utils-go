/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"example.com/user/string-utils/utils"
	"github.com/spf13/cobra"
)

var (
	debugMode           bool
	submatchFlag        bool
	lineNumberFlag      bool
	isInputPathFileMode bool
	filePathList        []string
	outputFilePath      string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "string-utils",
	Version:           "0.1.0-beta.1",
	Short:             "A utility for manipulating strings inside files",
	Long:              `Contains common utility functions for maniupulating strings in mutliple files`,
	PersistentPreRun:  utils.PreRun(&debugMode),
	PersistentPostRun: utils.PostRun(),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.string-utils.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolVarP(&debugMode, "debug", "d", false, "Enable debug output")
	rootCmd.PersistentFlags().BoolVarP(&isInputPathFileMode, "filePath", "f", false, "Enable Input Path as file")
	rootCmd.PersistentFlags().StringVarP(&outputFilePath, "outputFile", "o", "", "Enable writing to outputFile")
}

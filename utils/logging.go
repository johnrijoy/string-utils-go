package utils

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	InfoLogger  *log.Logger
	DebugLogger *log.Logger
	ErrorLogger *log.Logger
	outputFile  *os.File
)

func init() {
	InfoLogger = log.New(os.Stdout, "", 0)
	DebugLogger = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func PreRun(debugMode *bool) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		if *debugMode {
			InfoLn("Debug logs enabled")
		} else {
			DebugLogger.SetOutput(io.Discard)
		}

		if cmd.Flag("outputFile").Changed {
			filePath, _ := cmd.Flags().GetString("outputFile")
			outputFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
			HandlePanic(err)

			InfoLn("Saving Output to: ", filePath)

			InfoLogger.SetOutput(outputFile)
		}
	}
}

func PostRun() func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		outputFile.Close()
	}
}

// Logging helper functions

func Info(v ...any) {
	InfoLogger.Print(v...)
}

func InfoLn(v ...any) {
	InfoLogger.Println(v...)
}

func Infof(format string, v ...any) {
	InfoLogger.Printf(format, v...)
}

func Debug(v ...any) {
	DebugLogger.Print(v...)
}

func DebugLn(v ...any) {
	DebugLogger.Println(v...)
}

func Debugf(format string, v ...any) {
	DebugLogger.Printf(format, v...)
}

func Error(v ...any) {
	ErrorLogger.Print(v...)
}

func ErrorLn(v ...any) {
	ErrorLogger.Println(v...)
}

func Errorf(format string, v ...any) {
	ErrorLogger.Printf(format, v...)
}

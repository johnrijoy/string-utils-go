package utils

import "os"

func HandlePanic(err error) {
	if err != nil {
		ErrorLn(err)
		panic(err)
	}
}

func HandleExit(err error) {
	if err != nil {
		ErrorLn(err)
		ErrorLn("Exiting ...")
		os.Exit(1)
	}
}

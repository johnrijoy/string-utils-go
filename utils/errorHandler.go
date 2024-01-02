package utils

import "fmt"

func HandlePanic(err error) {
	if err != nil {
		fmt.Println("Eoncountered error: ", err)
		panic(err)
	}
}

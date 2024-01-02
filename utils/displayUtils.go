package utils

import (
	"fmt"
)

func PrintOccurenceMap(resultList map[string][]string) {
	for key := range resultList {
		fmt.Println("file: ", key)
		PrintOccurenceList(resultList[key])
	}
}

func PrintSubmatchMap(resultList map[string][][]string) {
	for key := range resultList {
		fmt.Println("file: ", key)
		for _, matches := range resultList[key] {
			for _, match := range matches {
				fmt.Print(match + " ")
			}
			fmt.Println()
		}
	}
}

func PrintOccurenceList(resultList []string) {

	for _, occur := range resultList {
		fmt.Println(occur)
	}

}

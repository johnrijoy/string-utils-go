package utils

import "strings"

func PrintOccurenceMap(resultList map[string][]string) {
	for key := range resultList {
		InfoLn("file: ", key)
		PrintOccurenceList(resultList[key])
	}
}

func PrintSubmatchMap(resultList map[string][][]string) {
	for key := range resultList {
		InfoLn("file: ", key)
		for _, matches := range resultList[key] {
			fullLine := matches[0]
			subMatches := strings.Join(matches[1:], " ")

			InfoLn(fullLine)
			InfoLn(subMatches)
		}
	}
}

func PrintOccurenceList(resultList []string) {

	for _, occur := range resultList {
		InfoLn(occur)
	}

}

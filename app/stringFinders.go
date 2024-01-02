package app

import (
	"os"
	"regexp"

	"example.com/user/string-utils/utils"
)

func FindAllFromFile(regexExp string, filePath string) (occurrenceList []string) {

	file, err := os.ReadFile(filePath)
	utils.HandlePanic(err)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	allOccur := r.FindAll(file, -1)

	for _, occur := range allOccur {
		occurrenceList = append(occurrenceList, string(occur))
	}

	return
}

func FindAllFromGlobPattern(regexExp string, globPattern string) (occurrenceMap map[string][]string) {

	fileNames := utils.FindFilesFromGlobPattern(globPattern)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap = make(map[string][]string)

	for _, fileName := range fileNames {

		var occurrenceList []string

		file, err := os.ReadFile(fileName)
		utils.HandlePanic(err)

		occurrenceList = findMatches(r, file, occurrenceList)

		if len(occurrenceList) != 0 {
			occurrenceMap[fileName] = occurrenceList
		}
	}

	return
}

func FindAllSubmatchFromGlobPattern(regexExp string, globPattern string) map[string][][]string {

	fileNames := utils.FindFilesFromGlobPattern(globPattern)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap := make(map[string][][]string)

	for _, fileName := range fileNames {

		var occurrenceList [][]string

		file, err := os.ReadFile(fileName)
		utils.HandlePanic(err)

		occurrenceList = findSubmatches(r, file, occurrenceList)
		if len(occurrenceList) != 0 {
			occurrenceMap[fileName] = occurrenceList
		}
	}

	return occurrenceMap
}

func findMatches(r *regexp.Regexp, file []byte, occurrenceList []string) []string {
	allOccur := r.FindAll(file, -1)

	for _, occur := range allOccur {
		occurrenceList = append(occurrenceList, string(occur))
	}
	return occurrenceList
}

func findSubmatches(r *regexp.Regexp, file []byte, occurrenceList [][]string) [][]string {
	allMatches := r.FindAllSubmatch(file, -1)

	for _, match := range allMatches {
		var matchList []string
		for _, occur := range match {
			matchList = append(matchList, string(occur))
		}
		occurrenceList = append(occurrenceList, matchList)
	}
	return occurrenceList
}

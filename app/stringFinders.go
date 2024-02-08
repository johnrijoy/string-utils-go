package app

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"example.com/user/string-utils/utils"
)

// Standard matches

func FindAllFromFile(regexExp string, filePath string) (occurrenceList []string) {

	file, err := os.ReadFile(filePath)
	utils.HandlePanic(err)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceList = findMatches(r, file, occurrenceList)

	return
}

func FindAllFromGlobPattern(regexExp string, globPatterns []string) (occurrenceMap map[string][]string) {

	fileNamesList := utils.FindFilesFromGlobPatterns(globPatterns)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap = make(map[string][]string)

	for _, fileName := range fileNamesList {

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

func FindAllLinesFromGlobPattern(regexExp string, globPatterns []string) (occurrenceMap map[string][]string) {

	fileNamesList := utils.FindFilesFromGlobPatterns(globPatterns)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap = make(map[string][]string)

	for _, fileName := range fileNamesList {

		var occurrenceList []string

		file, err := os.Open(fileName)
		utils.HandlePanic(err)

		occurrenceList = findLineMatches(r, file, occurrenceList)

		if len(occurrenceList) != 0 {
			occurrenceMap[fileName] = occurrenceList
		}
	}

	return
}

// Submatches

func FindAllSubmatchFromGlobPattern(regexExp string, globPatterns []string) map[string][][]string {

	fileNamesList := utils.FindFilesFromGlobPatterns(globPatterns)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap := make(map[string][][]string)

	for _, fileName := range fileNamesList {

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

func FindAllLinesSubmatchFromGlobPattern(regexExp string, globPatterns []string) map[string][][]string {

	fileNamesList := utils.FindFilesFromGlobPatterns(globPatterns)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap := make(map[string][][]string)

	for _, fileName := range fileNamesList {

		var occurrenceList [][]string

		file, err := os.Open(fileName)
		utils.HandlePanic(err)

		occurrenceList = findLineSubmatches(r, file, occurrenceList)
		if len(occurrenceList) != 0 {
			occurrenceMap[fileName] = occurrenceList
		}
	}

	return occurrenceMap
}

// Utils
func findMatches(r *regexp.Regexp, file []byte, occurrenceList []string) []string {
	allOccur := r.FindAll(file, -1)

	for _, occur := range allOccur {
		occurrenceList = append(occurrenceList, string(occur))
	}
	return occurrenceList
}

func findLineMatches(r *regexp.Regexp, file *os.File, occurrenceList []string) []string {
	fileScanner := bufio.NewScanner(file)

	line := 1
	for fileScanner.Scan() {
		allOccur := r.FindAll(fileScanner.Bytes(), -1)

		for _, occur := range allOccur {
			modifOccur := fmt.Sprintf("%2d: %s", line, string(occur))
			occurrenceList = append(occurrenceList, modifOccur)
		}
		line++
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

func findLineSubmatches(r *regexp.Regexp, file *os.File, occurrenceList [][]string) [][]string {
	fileScanner := bufio.NewScanner(file)

	line := 1
	for fileScanner.Scan() {
		allMatches := r.FindAllSubmatch(fileScanner.Bytes(), -1)

		for _, match := range allMatches {
			var matchList []string
			for _, occur := range match {
				matchList = append(matchList, string(occur))
			}

			fullLine := matchList[0]
			modifLine := fmt.Sprintf("%2d: %s", line, fullLine)
			matchList[0] = modifLine
			occurrenceList = append(occurrenceList, matchList)

		}
		line++
	}
	return occurrenceList
}

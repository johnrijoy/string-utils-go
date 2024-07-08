package app

import (
	"os"
	"regexp"
	"strconv"

	"github.com/johnrijoy/string-utils-go/utils"
)

type ChangeResult struct {
	before, after string
}

func ReplaceAllInFile(regexExp, filePath, replaceText string) (occurrenceList MatchResult) {

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	file, err := os.ReadFile(filePath)
	utils.HandlePanic(err)

	allOccur := r.FindAll(file, -1)

	for _, occur := range allOccur {
		occurrenceList.Matches = append(occurrenceList.Matches, string(occur))
	}

	// to convert escape characters
	replaceText, err = strconv.Unquote(`"` + replaceText + `"`)
	utils.HandlePanic(err)

	newFile := r.ReplaceAll(file, []byte(replaceText))

	err = os.WriteFile(filePath, newFile, 0644)
	utils.HandlePanic(err)

	return
}

func ReplaceAllInGlobPattern(regexExp, globPattern, replaceText string, templateMode bool) (occurrenceMap map[string]MatchResult) {

	fileNames := utils.FindFilesFromGlobPattern(globPattern)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap = make(map[string]MatchResult)

	for _, fileName := range fileNames {

		var newFile []byte

		file, err := os.ReadFile(fileName)
		utils.HandlePanic(err)

		if templateMode {
			replaceTemplate, err := os.ReadFile(replaceText)
			utils.HandlePanic(err)
			newFile = r.ReplaceAllLiteral(file, replaceTemplate)
		} else {
			// to convert escape characters
			replaceText, err = strconv.Unquote(`"` + replaceText + `"`)
			utils.HandlePanic(err)

			newFile = r.ReplaceAllLiteral(file, []byte(replaceText))
		}

		err = os.WriteFile(fileName, newFile, 0644)
		utils.HandlePanic(err)

		// collect changes
		occurrenceList := findMatches(r, file)

		if occurrenceList.GetCount() != 0 {
			occurrenceMap[fileName] = occurrenceList
		}
	}

	return
}

func ReplaceAllSubmatchesInGlobPattern(regexExp, globPattern, replaceText string, templateMode bool) map[string][]ChangeResult {

	fileNames := utils.FindFilesFromGlobPattern(globPattern)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	changeMap := make(map[string][]ChangeResult)

	for _, fileName := range fileNames {
		var newFile []byte

		file, err := os.ReadFile(fileName)
		utils.HandlePanic(err)

		if templateMode {
			replaceTemplate, err := os.ReadFile(replaceText)
			utils.HandlePanic(err)
			newFile = r.ReplaceAll(file, replaceTemplate)
		} else {
			// to convert escape characters
			replaceText, err = strconv.Unquote(`"` + replaceText + `"`)
			utils.HandlePanic(err)

			newFile = r.ReplaceAll(file, []byte(replaceText))
		}

		// get list of changes with replacement
		changeList := getChangeList(r, file, newFile)
		if len(changeList) != 0 {
			changeMap[fileName] = changeList
		}

		err = os.WriteFile(fileName, newFile, 0644)
		utils.HandlePanic(err)
	}

	return changeMap
}

func ReplaceAllInGlobPatterns(regexExp string, globPatterns []string, replaceText string, templateMode bool) (occurrenceMap map[string]MatchResult) {

	fileNames := utils.FindFilesFromGlobPatterns(globPatterns)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap = make(map[string]MatchResult)

	for _, fileName := range fileNames {

		var newFile []byte

		file, err := os.ReadFile(fileName)
		utils.HandlePanic(err)

		occurrenceList := findMatches(r, file)

		if occurrenceList.GetCount() != 0 {
			occurrenceMap[fileName] = occurrenceList
		}

		if templateMode {
			replaceTemplate, err := os.ReadFile(replaceText)
			utils.HandlePanic(err)
			newFile = r.ReplaceAllLiteral(file, replaceTemplate)
		} else {
			// to convert escape characters
			replaceText, err = strconv.Unquote(`"` + replaceText + `"`)
			utils.HandlePanic(err)

			newFile = r.ReplaceAllLiteral(file, []byte(replaceText))
		}

		err = os.WriteFile(fileName, newFile, 0644)
		utils.HandlePanic(err)
	}

	return
}

func ReplaceAllSubmatchesInGlobPatterns(regexExp string, globPatterns []string, replaceText string, templateMode bool) map[string][]ChangeResult {

	fileNames := utils.FindFilesFromGlobPatterns(globPatterns)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	changeMap := make(map[string][]ChangeResult)

	for _, fileName := range fileNames {

		var newFile []byte

		file, err := os.ReadFile(fileName)
		utils.HandlePanic(err)

		if templateMode {
			replaceTemplate, err := os.ReadFile(replaceText)
			utils.HandlePanic(err)
			newFile = r.ReplaceAll(file, replaceTemplate)
		} else {
			// to convert escape characters
			replaceText, err = strconv.Unquote(`"` + replaceText + `"`)
			utils.HandlePanic(err)

			newFile = r.ReplaceAll(file, []byte(replaceText))
		}

		// get list of changes with replacement
		changeList := getChangeList(r, file, newFile)
		if len(changeList) != 0 {
			changeMap[fileName] = changeList
		}

		err = os.WriteFile(fileName, newFile, 0644)
		utils.HandlePanic(err)
	}

	return changeMap
}

// Helpers

func getChangeList(r *regexp.Regexp, file []byte, newFile []byte) (changeList []ChangeResult) {
	for _, match := range r.FindAllSubmatchIndex(file, -1) {
		beforeValue := string(file[match[0]:match[1]])
		afterValue := string(newFile[match[0]:match[1]])

		changeList = append(changeList, ChangeResult{beforeValue, afterValue})

	}
	return changeList
}

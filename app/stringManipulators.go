package app

import (
	"os"
	"regexp"
	"strconv"

	"example.com/user/string-utils/utils"
)

func ReplaceAllInFile(regexExp, filePath, replaceText string) (occurrenceList []string) {

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	file, err := os.ReadFile(filePath)
	utils.HandlePanic(err)

	allOccur := r.FindAll(file, -1)

	for _, occur := range allOccur {
		occurrenceList = append(occurrenceList, string(occur))
	}

	// to convert escape characters
	replaceText, err = strconv.Unquote(`"` + replaceText + `"`)
	utils.HandlePanic(err)

	newFile := r.ReplaceAll(file, []byte(replaceText))

	err = os.WriteFile(filePath, newFile, 0644)
	utils.HandlePanic(err)

	return
}

func ReplaceAllInGlobPattern(regexExp, globPattern, replaceText string, templateMode bool) (occurrenceMap map[string][]string) {

	fileNames := utils.FindFilesFromGlobPattern(globPattern)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap = make(map[string][]string)

	for _, fileName := range fileNames {

		var occurrenceList []string
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
		occurrenceList = findMatches(r, file, occurrenceList)

		if len(occurrenceList) != 0 {
			occurrenceMap[fileName] = occurrenceList
		}
	}

	return
}

func ReplaceAllSubmatchesInGlobPattern(regexExp, globPattern, replaceText string, templateMode bool) map[string][][]string {

	fileNames := utils.FindFilesFromGlobPattern(globPattern)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap := make(map[string][][]string)

	for _, fileName := range fileNames {

		var occurrenceList [][]string
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
		occurrenceList = getChangeList(r, file, newFile, occurrenceList)
		if len(occurrenceList) != 0 {
			occurrenceMap[fileName] = occurrenceList
		}

		err = os.WriteFile(fileName, newFile, 0644)
		utils.HandlePanic(err)
	}

	return occurrenceMap
}

func ReplaceAllInGlobPatterns(regexExp string, globPatterns []string, replaceText string, templateMode bool) (occurrenceMap map[string][]string) {

	fileNames := utils.FindFilesFromGlobPatterns(globPatterns)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap = make(map[string][]string)

	for _, fileName := range fileNames {

		var occurrenceList []string
		var newFile []byte

		file, err := os.ReadFile(fileName)
		utils.HandlePanic(err)

		occurrenceList = findMatches(r, file, occurrenceList)

		if len(occurrenceList) != 0 {
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

func ReplaceAllSubmatchesInGlobPatterns(regexExp string, globPatterns []string, replaceText string, templateMode bool) map[string][][]string {

	fileNames := utils.FindFilesFromGlobPatterns(globPatterns)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap := make(map[string][][]string)

	for _, fileName := range fileNames {

		var occurrenceList [][]string
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
		occurrenceList = getChangeList(r, file, newFile, occurrenceList)
		if len(occurrenceList) != 0 {
			occurrenceMap[fileName] = occurrenceList
		}

		err = os.WriteFile(fileName, newFile, 0644)
		utils.HandlePanic(err)
	}

	return occurrenceMap
}

func getChangeList(r *regexp.Regexp, file []byte, newFile []byte, occurrenceList [][]string) [][]string {
	for _, match := range r.FindAllSubmatchIndex(file, -1) {
		beforeValue := string(file[match[0]:match[1]])
		afterValue := string(newFile[match[0]:match[1]])

		occurrenceList = append(occurrenceList, []string{beforeValue, afterValue})

	}
	return occurrenceList
}

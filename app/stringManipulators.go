package app

import (
	"os"
	"regexp"
	"strconv"

	"example.com/user/string-utils/utils"
)

func ResplaceAllInFile(regexExp, filePath, replaceText string) (occurrenceList []string) {

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

func ResplaceAllInGlobPattern(regexExp, globPattern, replaceText string) (occurrenceMap map[string][]string) {

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

		// to convert escape characters
		replaceText, err = strconv.Unquote(`"` + replaceText + `"`)
		utils.HandlePanic(err)

		newFile := r.ReplaceAllLiteral(file, []byte(replaceText))

		err = os.WriteFile(fileName, newFile, 0644)
		utils.HandlePanic(err)
	}

	return
}

func ResplaceAllSubmatchesInGlobPattern(regexExp, globPattern, replaceText string) map[string][][]string {

	fileNames := utils.FindFilesFromGlobPattern(globPattern)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap := make(map[string][][]string)

	for _, fileName := range fileNames {

		var occurrenceList [][]string

		file, err := os.ReadFile(fileName)
		utils.HandlePanic(err)

		// to convert escape characters
		replaceText, err = strconv.Unquote(`"` + replaceText + `"`)
		utils.HandlePanic(err)

		newFile := r.ReplaceAll(file, []byte(replaceText))

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

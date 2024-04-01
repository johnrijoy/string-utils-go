package app

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sync"

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
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	for _, fileName := range fileNamesList {
		wg.Add(1)

		go func(fileName string) {
			var occurrenceList []string

			file, err := os.ReadFile(fileName)
			utils.HandlePanic(err)

			occurrenceList = findMatches(r, file, occurrenceList)

			if len(occurrenceList) != 0 {
				mutex.Lock()
				occurrenceMap[fileName] = occurrenceList
				mutex.Unlock()
			}
			wg.Done()
		}(fileName)
	}

	wg.Wait()
	return
}

func FindAllLinesFromGlobPattern(regexExp string, globPatterns []string) (occurrenceMap map[string][]string) {

	fileNamesList := utils.FindFilesFromGlobPatterns(globPatterns)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap = make(map[string][]string)
	var waitGroup sync.WaitGroup
	var mutex sync.Mutex

	for _, fileN := range fileNamesList {
		waitGroup.Add(1)

		go func(fileName string) {
			var occurrenceList []string

			file, err := os.ReadFile(fileName)
			utils.HandlePanic(err)

			occurrenceList = findLineMatches(r, file, occurrenceList)

			if len(occurrenceList) != 0 {
				mutex.Lock()
				occurrenceMap[fileName] = occurrenceList
				mutex.Unlock()
			}
			waitGroup.Done()
		}(fileN)
	}
	waitGroup.Wait()
	return
}

// Submatches

func FindAllSubmatchFromGlobPattern(regexExp string, globPatterns []string) map[string][][]string {

	fileNamesList := utils.FindFilesFromGlobPatterns(globPatterns)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap := make(map[string][][]string)
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	for _, fileName := range fileNamesList {
		wg.Add(1)

		go func(fileName string) {
			var occurrenceList [][]string

			file, err := os.ReadFile(fileName)
			utils.HandlePanic(err)

			occurrenceList = findSubmatches(r, file, occurrenceList)
			if len(occurrenceList) != 0 {
				mutex.Lock()
				occurrenceMap[fileName] = occurrenceList
				mutex.Unlock()
			}
			wg.Done()
		}(fileName)
	}
	wg.Wait()
	return occurrenceMap
}

func FindAllLinesSubmatchFromGlobPattern(regexExp string, globPatterns []string) map[string][][]string {

	fileNamesList := utils.FindFilesFromGlobPatterns(globPatterns)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap := make(map[string][][]string)

	for _, fileName := range fileNamesList {

		var occurrenceList [][]string

		file, err := os.ReadFile(fileName)
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

func findLineMatches(r *regexp.Regexp, file []byte, occurrenceList []string) []string {
	allLoc := r.FindAllIndex(file, -1)
	const newLineByte = '\n'
	for _, loc := range allLoc {
		utils.DebugLn("Index: ", loc)
		if len(loc) == 2 {
			count := bytes.Count(file[:loc[0]], []byte{newLineByte})
			modifOccur := fmt.Sprintf("%2d: %s", count+1, string(file[loc[0]:loc[1]]))
			occurrenceList = append(occurrenceList, modifOccur)
		}
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

func findLineSubmatches(r *regexp.Regexp, file []byte, occurrenceList [][]string) [][]string {
	allLoc := r.FindAllSubmatchIndex(file, -1)
	const newLineByte = '\n'

	for _, loc := range allLoc {

		utils.DebugLn("Index: ", loc)
		if len(loc) > 0 && len(loc)%2 == 0 {

			count := bytes.Count(file[:loc[0]], []byte{newLineByte})
			firstOccur := fmt.Sprintf("%2d: %s", count+1, string(file[loc[0]:loc[1]]))

			var matchList []string
			matchList = append(matchList, firstOccur)

			for i := 2; i < len(loc); i += 2 {
				matchList = append(matchList, string(file[loc[i]:loc[i+1]]))
			}

			occurrenceList = append(occurrenceList, matchList)
		}
	}

	return occurrenceList
}

func findAllHelper[T []any](regexExp string, globPatterns []string, matcherFinder func(r *regexp.Regexp, file []byte, occurrenceList T) T) map[string]T {
	fileNamesList := utils.FindFilesFromGlobPatterns(globPatterns)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceMap := make(map[string]T)
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	for _, fileName := range fileNamesList {
		wg.Add(1)

		go func(fileName string) {
			var occurrenceList T

			file, err := os.ReadFile(fileName)
			utils.HandlePanic(err)

			occurrenceList = matcherFinder(r, file, occurrenceList)

			if len(occurrenceList) != 0 {
				mutex.Lock()
				occurrenceMap[fileName] = occurrenceList
				mutex.Unlock()
			}
			wg.Done()
		}(fileName)
	}

	wg.Wait()

	return occurrenceMap
}

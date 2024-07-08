package app

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sync"

	"github.com/johnrijoy/string-utils-go/utils"
)

type Matches interface {
	GetCount() int
}

type MatchResult struct {
	Matches []string
}

func (res MatchResult) GetCount() int {
	return len(res.Matches)
}

type SubMatchResult struct {
	Matches [][]string
}

func (res SubMatchResult) GetCount() int {
	return len(res.Matches)
}

// Standard matches

func FindAllFromFile(regexExp string, filePath string) (occurrenceList MatchResult) {

	file, err := os.ReadFile(filePath)
	utils.HandlePanic(err)

	r, err := regexp.Compile(regexExp)
	utils.HandlePanic(err)

	occurrenceList = findMatches(r, file)

	return
}

func FindAllFromGlobPattern(regexExp string, globPatterns []string) map[string]MatchResult {

	return findAllHelper(regexExp, globPatterns, findMatches)
}

func FindAllLinesFromGlobPattern(regexExp string, globPatterns []string) (occurrenceMap map[string]MatchResult) {

	return findAllHelper(regexExp, globPatterns, findLineMatches)
}

// Submatches

func FindAllSubmatchFromGlobPattern(regexExp string, globPatterns []string) map[string]SubMatchResult {

	return findAllHelper(regexExp, globPatterns, findSubmatches)
}

func FindAllLinesSubmatchFromGlobPattern(regexExp string, globPatterns []string) map[string]SubMatchResult {

	return findAllHelper(regexExp, globPatterns, findLineSubmatches)
}

// Utils
func findMatches(r *regexp.Regexp, file []byte) (occurrenceList MatchResult) {
	allOccur := r.FindAll(file, -1)

	for _, occur := range allOccur {
		occurrenceList.Matches = append(occurrenceList.Matches, string(occur))
	}
	return occurrenceList
}

func findLineMatches(r *regexp.Regexp, file []byte) (occurrenceList MatchResult) {
	allLoc := r.FindAllIndex(file, -1)
	const newLineByte = '\n'
	for _, loc := range allLoc {
		utils.DebugLn("Index: ", loc)
		if len(loc) == 2 {
			count := bytes.Count(file[:loc[0]], []byte{newLineByte})
			modifOccur := fmt.Sprintf("%2d: %s", count+1, string(file[loc[0]:loc[1]]))
			occurrenceList.Matches = append(occurrenceList.Matches, modifOccur)
		}
	}
	return occurrenceList
}

func findSubmatches(r *regexp.Regexp, file []byte) (occurrenceList SubMatchResult) {
	allMatches := r.FindAllSubmatch(file, -1)

	for _, match := range allMatches {
		var matchList []string
		for _, occur := range match {
			matchList = append(matchList, string(occur))
		}
		occurrenceList.Matches = append(occurrenceList.Matches, matchList)
	}
	return occurrenceList
}

func findLineSubmatches(r *regexp.Regexp, file []byte) (occurrenceList SubMatchResult) {
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

			occurrenceList.Matches = append(occurrenceList.Matches, matchList)
		}
	}

	return occurrenceList
}

func findAllHelper[T Matches](regexExp string, globPatterns []string, matcherFinder func(r *regexp.Regexp, file []byte) T) map[string]T {
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

			occurrenceList = matcherFinder(r, file)

			if occurrenceList.GetCount() != 0 {
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

package app

import (
	"bufio"
	"os"
	"strings"

	"github.com/johnrijoy/string-utils-go/utils"
)

func PrintOccurenceMap(resultList map[string]MatchResult) {
	for key := range resultList {
		utils.InfoLn("file: ", key)
		PrintOccurenceList(resultList[key])
	}
}

func PrintSubmatchMap(resultList map[string]SubMatchResult) {
	for key := range resultList {
		utils.InfoLn("file: ", key)
		for _, matches := range resultList[key].Matches {
			fullLine := matches[0]
			subMatches := strings.Join(matches[1:], " ")

			utils.InfoLn(fullLine)
			utils.InfoLn(subMatches)
		}
	}
}

func PrintOccurenceList(resultList MatchResult) {

	for _, occur := range resultList.Matches {
		utils.InfoLn(occur)
	}

}

func PrintChangeMap(resultList map[string][]ChangeResult) {
	for key := range resultList {
		utils.InfoLn("file: ", key)
		for _, change := range resultList[key] {

			utils.InfoLn("Before:")
			utils.InfoLn(change.before)
			utils.InfoLn("After:")
			utils.InfoLn(change.after)
		}
	}
}

func PrintRenameMap(renameMap map[string]string) {
	utils.InfoLn("\nExpected output:")
	for fileName := range renameMap {
		utils.InfoLn(fileName, "->", renameMap[fileName])
	}
}

func TestPrompt(label string, defaultVal bool) bool {
	testResult := defaultVal
	r := bufio.NewReader(os.Stdin)

	cmdMsg := ""
	if defaultVal {
		cmdMsg = " ( [y]/n ):"
	} else {
		cmdMsg = " ( y/[n] ):"
	}

	utils.Info(label + cmdMsg)
	s, _ := r.ReadString('\n')
	if s != "" {
		switch strings.TrimSpace(s) {
		case "Y", "y", "Yes", "yes":
			testResult = true
		case "N", "n", "No", "no":
			testResult = false
		default:
			testResult = defaultVal
		}
	}

	return testResult
}

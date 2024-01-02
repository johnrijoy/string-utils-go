package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar"
)

func FindFilesFromGlobPattern(globPattern string) (fileList []string) {

	fileList, err := doublestar.Glob(globPattern)
	HandlePanic(err)

	fileList = filterFiles(fileList)

	return
}

func FindFilesFromPathPattern(pathPattern string) (fileList []string) {

	rootPath := strings.SplitN(pathPattern, "*", 2)[0]

	fmt.Println("Root Path: ", rootPath)

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path, " -- ", info.Name())

		// BUG: Match does support ** patterns
		match, err := filepath.Match(pathPattern, path)
		HandlePanic(err)

		if match {
			fileList = append(fileList, info.Name())
		}

		return nil
	})
	HandlePanic(err)

	return
}

func filterFiles(pathList []string) (resultList []string) {
	for _, pathName := range pathList {
		fileInfo, err := os.Stat(pathName)
		HandlePanic(err)

		if !fileInfo.IsDir() {
			resultList = append(resultList, pathName)
		}
	}

	return
}

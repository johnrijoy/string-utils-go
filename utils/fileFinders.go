package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar"
)

func FindFilesFromGlobPattern(globPattern string) (fileList []string) {

	DebugLn("GlobPattern: ", globPattern)
	fileList, err := doublestar.Glob(globPattern)
	HandlePanic(err)

	DebugLn("Files: ", fileList)

	fileList = filterFiles(fileList)

	return
}

func FindFilesFromGlobPatterns(globPatterns []string) (fileList []string) {

	for _, glglobPattern := range globPatterns {
		fileList = append(fileList, FindFilesFromGlobPattern(glglobPattern)...)
	}

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

func ReadInputPathFile(inputFile string) []string {
	var filePathList []string

	file, err := os.Open(inputFile)
	HandlePanic(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		filePathList = append(filePathList, fileScanner.Text())
	}

	return filePathList
}

func filterFiles(pathList []string) (resultList []string) {
	for _, pathName := range pathList {
		fileInfo, err := os.Lstat(pathName)
		HandlePanic(err)

		// check if file is not a directory
		if !fileInfo.IsDir() {
			// Check if file is a symlink, will return error if not
			if _, err = os.Readlink(pathName); err != nil {
				resultList = append(resultList, pathName)
			}
		}
	}

	return
}

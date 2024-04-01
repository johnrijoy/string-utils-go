package app

import (
	"os"
	"path/filepath"
	"regexp"

	"example.com/user/string-utils/utils"
)

func RenameFiles(basePath, filePattern, newFileNamePattern string) error {
	fileList, err := utils.FindFilesInDir(basePath)
	if err != nil {
		return err
	}

	r, err := regexp.Compile(filePattern)
	utils.HandlePanic(err)

	renameMap := make(map[string]string)

	for _, fileName := range fileList {
		if r.MatchString(fileName) {
			newFileName := r.ReplaceAllString(fileName, newFileNamePattern)
			renameMap[fileName] = newFileName
		}
	}

	PrintRenameMap(renameMap)

	if !TestPrompt("Do you want to continue with the rename?", false) {
		utils.InfoLn("Rename cancelled...")
		return nil
	}

	for fileName := range renameMap {
		if err := RenameFile(basePath, fileName, renameMap[fileName]); err != nil {
			return err
		}
	}
	utils.InfoLn("Rename successful...")
	return nil
}

func RenameFile(basePath, fileName, newFileName string) error {
	filePathAbs := filepath.Join(basePath, fileName)
	newFilePathAbs := filepath.Join(basePath, newFileName)
	return os.Rename(filePathAbs, newFilePathAbs)
}

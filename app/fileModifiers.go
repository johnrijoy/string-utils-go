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

	for _, fileName := range fileList {
		newFileName := r.ReplaceAllString(fileName, newFileNamePattern)
		if err := RenameFile(basePath, fileName, newFileName); err != nil {
			return err
		}

	}

	return nil
}

func RenameFile(basePath, fileName, newFileName string) error {
	filePathAbs := filepath.Join(basePath, fileName)
	newFilePathAbs := filepath.Join(basePath, newFileName)
	return os.Rename(filePathAbs, newFilePathAbs)
}

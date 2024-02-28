package generator

import (
	"os"
	"path/filepath"
)

func GetFiles(path string) ([]string, error) {

	var filesPaths []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && filepath.Base(path)[0] == '.' {
			return filepath.SkipDir
		}
		if info.IsDir() {
			return nil
		}

		filesPaths = append(filesPaths, path)

		return nil
	})

	if err != nil {
		return []string{}, err
	}

	return filesPaths, nil

}

func WriteFile(file string, content []byte) error {

	// if destion dir not exist create it
	dir := filepath.Dir(file)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	err := os.WriteFile(file, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetFileNewName(file string, templateNewFileNames []interface{}) string {
	var newFile string
	fileName := filepath.Base(file)

	// remove file extension
	fileName = fileName[:len(fileName)-len(filepath.Ext(fileName))]

	for _, variableVal := range templateNewFileNames {
		if fileName == variableVal.(map[string]interface{})["from"].(string) {
			newFile = variableVal.(map[string]interface{})["to"].(string)
		} else {
			newFile = fileName
		}
	}
	return newFile + filepath.Ext(file)
}

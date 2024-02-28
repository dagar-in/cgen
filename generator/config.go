package generator

import (
	"encoding/json"
	"os"
)

func GetConfig(confPath string) (any, string, error) {

	wd, err := os.Getwd()
	if err != nil {
		return nil, "", err
	}

	configPathToUse := wd + "/.gen.config.json"
	if confPath != "" {
		configPathToUse = confPath
	}

	if _, err := os.Stat(configPathToUse); os.IsNotExist(err) {
		return nil, wd, err
	}

	content, err := os.ReadFile(configPathToUse)
	if err != nil {
		return nil, wd, err
	}

	if !json.Valid(content) {
		return nil, wd, err
	}

	obj := make(map[string]interface{})
	err = json.Unmarshal(content, &obj)
	if err != nil {
		return nil, wd, err
	}

	return obj, wd, nil

}

package generator

import (
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
)

var config string = ""

var gen = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen", "g"},
	Short:   "Generate codebase from a template file/dir",

	Run: func(cmd *cobra.Command, args []string) {

		conf, pwd, err := GetConfig(config)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		templatesConf := conf.(map[string]interface{})["templates"].(map[string]interface{})
		templatePath := filepath.Join(pwd, templatesConf["path"].(string))
		destination := filepath.Join(pwd, templatesConf["destination"].(string))
		templateVariables := templatesConf["variables"].([]interface{})
		templateNewFileNames := templatesConf["fileRenames"].([]interface{})

		files, err := GetFiles(templatePath)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		if len(files) == 0 {
			fmt.Println("No files found in the template folder")
			return
		}

		parsedVariables := make(map[string]string)

		for _, value := range templateVariables {
			v := value.(map[string]interface{})
			parsedVariables[v["name"].(string)] = v["value"].(string)
		}

		for _, file := range files {
			content, err := ParseFile(file, parsedVariables)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			newFileName := GetFileNewName(file, templateNewFileNames)

			destinationFilPath, _ := filepath.Rel(templatePath, file)
			destinationLocation := filepath.Join(destination, destinationFilPath)

			fmt.Println("New file name: ", filepath.Join(destinationLocation, newFileName))
			err = WriteFile(filepath.Join(destinationLocation, newFileName), content)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
		}
	},
}

func Init(rootCmd *cobra.Command) {

	gen.Flags().StringVarP(&config, "config", "c", "", "Path to the config file")

	rootCmd.AddCommand(gen)
}

package generator

import (
	"bytes"
	"fmt"
	"text/template"
)

func ParseFile(file string, variables map[string]string) ([]byte, error) {

	fmt.Println("Parsing file: ", file)
	templateFil, err := template.ParseFiles(file)
	if err != nil {
		return nil, err
	}
	var resultFile bytes.Buffer
	if err := templateFil.Execute(&resultFile, variables); err != nil {
		return nil, err
	}

	return resultFile.Bytes(), nil
}

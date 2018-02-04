package util

import (
	"io/ioutil"
	"strings"
)

func OpenChecklist(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	fileContents := string(data)
	fileLines := strings.Split(fileContents, "\n")
	return fileLines, nil
}
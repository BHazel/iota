package util

import (
	"github.com/bhazel/iota/config"	
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func SetChecklistFilename(filename string) string {
	if strings.LastIndex(filename, ".") != -1 {
		return filename
	} else {
		return fmt.Sprintf("%s.%s", filename, config.ChecklistFileExtension)
	}
}

func FileExists(filename string) (bool, error) {
	if _, err := os.Stat(filename); err == nil {
		return true, nil
	} else if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false, err
	} else {
		return false, err
	}
}

func OpenChecklistFile(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	fileContents := string(data)
	return strings.Split(fileContents, "\n"), nil
}
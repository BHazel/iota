package util

import (
	"github.com/bhazel/iota/config"
	"fmt"
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
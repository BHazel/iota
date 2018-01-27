package util

import (
	"github.com/bhazel/iota/config"
	"fmt"
	"strings"
)

func SetChecklistFilename(filename string) string {
	if strings.LastIndex(filename, ".") != -1 {
		return filename
	} else {
		return fmt.Sprintf("%s.%s", filename, config.ChecklistFileExtension)
	}
}
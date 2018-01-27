package util

import (
	"fmt"
	"os"
)

func PrintErr(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
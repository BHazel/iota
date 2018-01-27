package cmd

import (
	"fmt"
	"os"
	"github.com/bhazel/iota/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use: "init",
	Short: "Creates a new checklist",
	Long: "Creates a new checklist with specified name",
	Run: func(cmd *cobra.Command, args []string) {
		var filename string
		if len(args) == 0 {
			filename = setChecklistFilename("checklist")
		} else {
			filename = setChecklistFilename(args[0])
		}

		if _, err := os.Stat(filename); err == nil {
			fmt.Fprintf(os.Stderr, config.InitErrorFileExists, filename)
			os.Exit(config.EXIT_FILE_EXISTS)
		}

		_, err := os.Create(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, config.GenericError, err)
			os.Exit(config.EXIT_FILE_CREATE_ERROR)
		}

		fmt.Printf(config.InitFileCreated, filename)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func setChecklistFilename(filename string) string {
	return fmt.Sprintf("%s.%s", filename, config.ChecklistFileExtension)
}
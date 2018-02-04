package cmd

import (
	"fmt"
	"os"
	"github.com/bhazel/iota/cli/util"
	"github.com/bhazel/iota/config"
	"github.com/spf13/cobra"
)

var (
	loadCmd = &cobra.Command{
		Use: "load",
		Short: "Loads a checklist file",
		Long: "Loads a checklist file",
		Run: func(cmd *cobra.Command, args []string) {
			var filename string
			if len(args) == 0 {
				filename = util.SetChecklistFilename("checklist")
			} else {
				filename = util.SetChecklistFilename(args[0])
			}

			exists, err := util.FileExists(filename)
			if err != nil {
				util.PrintErr(config.GenericError, err)
				os.Exit(config.EXIT_FAILURE)
			} else if exists == false {
				util.PrintErr(config.ErrorFileNotFound, filename)
				os.Exit(config.EXIT_FILE_NOT_FOUND)
			}

			lines, err := util.OpenChecklistFile(filename)
			if err != nil {
				os.Exit(config.EXIT_FAILURE)
			}

			for i := range lines {
				fmt.Println(lines[i])
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(loadCmd)
}
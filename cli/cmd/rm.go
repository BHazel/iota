package cmd

import (
	"fmt"
	"os"
	"github.com/bhazel/iota/config"
	"github.com/bhazel/iota/cli/util"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use: "rm",
	Short: "Delete a checklist",
	Long: "Permenantly delete a checklist file",
	Run: func(cmd *cobra.Command, args []string) {
		var filename string
		if len(args) == 0 {
			util.PrintErr(config.ErrorFileNotSpecified)
			os.Exit(config.EXIT_FILE_NOT_SPECIFIED)
		}

		filename = util.SetChecklistFilename(args[0])
		if _, err := os.Stat(filename); err == nil {
			os.Remove(filename)
		} else {
			util.PrintErr(config.ErrorFileNotFound, filename)
			os.Exit(config.EXIT_FILE_NOT_FOUND)
		}

		fmt.Printf(config.RmFileDeleted, filename)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
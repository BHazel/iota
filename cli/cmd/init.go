package cmd

import (
	"fmt"
	"os"
	"github.com/bhazel/iota/config"
	"github.com/bhazel/iota/cli/util"
	"github.com/spf13/cobra"
)

var (
	initForce bool
	initSamples bool
	initCmd = &cobra.Command{
		Use: "init",
		Short: "Creates a new checklist",
		Long: "Creates a new checklist with specified name",
		Run: func(cmd *cobra.Command, args []string) {
			var filename string
			if len(args) == 0 {
				filename = util.SetChecklistFilename("checklist")
			} else {
				filename = util.SetChecklistFilename(args[0])
			}

			if initForce != true {
				if _, err := util.FileExists(filename); err == nil {
					util.PrintErr(config.InitErrorFileExists, filename)
					os.Exit(config.EXIT_FILE_EXISTS)
				}
			}

			file, err := os.Create(filename)
			if err != nil {
				util.PrintErr(config.GenericError, err)
				os.Exit(config.EXIT_FILE_CREATE_ERROR)
			}

			fmt.Printf(config.InitFileCreated, filename)

			if initSamples == true {
				file.WriteString(getSamples())
			}

			file.Close()
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&initForce, "force", "f", false, "Force create the checklist, overwriting an existing one")
	initCmd.Flags().BoolVarP(&initSamples, "include-samples", "s", false, "Include checklist item samples")
}

func getSamples() string {
	return fmt.Sprintf("%s\n%s\n", config.InitSampleCheckedItem, config.InitSampleUncheckedItem)
}
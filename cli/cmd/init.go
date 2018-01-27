package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use: "init",
	Short: "Creates a new checklist",
	Long: "Creates a new checklist with specified name",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Init")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
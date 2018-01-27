package cmd

import (
	"fmt"
	"github.com/bhazel/iota/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Prints the Iota version",
	Long: "Prints the Iota version to standard output",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Iota v%s\n", config.CurrentVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
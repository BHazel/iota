package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "iota",
	Short: "Manages checklists",
	Long: "Manages plain-text checklists conforming to the Iota checklist specification",
}

func Execute() {
	rootCmd.Execute()
}
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	// Aliases: []string{"ver", "v"},
	Short: "Print the version number of Demo",
	Long:  `All software has versions. This is demo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Demo Version: v1.0")
	},
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thiago-scherrer/ocx/internal/tools"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configure the basic credentials",
	Long:  `Initializes the ocx and configures AWS access credentials and data.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		tools.SetupAWS()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}

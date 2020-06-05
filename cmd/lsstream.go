package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thiago-scherrer/ocx/internal/tools"
)

// lsstreamCmd represents the lsstream command
var lsstreamCmd = &cobra.Command{
	Use:   "lsstream",
	Short: "List log stream",
	Long:  `Show all log stream avaliable on cloudwath region account`,
	Run: func(cmd *cobra.Command, args []string) {
		tools.Lsstream()
	},
}

func init() {
	logCmd.AddCommand(lsstreamCmd)
}

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thiago-scherrer/ocx/internal/tools"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "List log groups",
	Long:  `Show all log groups available on cloudwath region account`,
	Run: func(cmd *cobra.Command, args []string) {
		c := tools.Client()
		tools.Group(c)
	},
}

func init() {
	logCmd.AddCommand(groupCmd)
}

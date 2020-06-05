package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thiago-scherrer/ocx/internal/tools"
)

// lsgroupCmd represents the lsgroup command
var lsgroupCmd = &cobra.Command{
	Use:   "lsgroup",
	Short: "List log groups",
	Long:  `Show all log groups avaliable on cloudwath region account`,
	Run: func(cmd *cobra.Command, args []string) {
		tools.LsGroup()
	},
}

func init() {
	logCmd.AddCommand(lsgroupCmd)
}

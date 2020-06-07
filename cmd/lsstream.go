package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thiago-scherrer/ocx/internal/tools"
)

// streamCmd represents the stream command
var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "List log stream",
	Long:  `Show all log stream avaliable on cloudwath region account`,
	Run: func(cmd *cobra.Command, args []string) {
		tools.Stream(args)
	},
}

func init() {
	logCmd.AddCommand(streamCmd)
}

package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/thiago-scherrer/ocx/internal/tools"
)

// streamCmd represents the stream command
var streamCmd = &cobra.Command{
	Use:   "stream --gname GROUP_NAME",
	Short: "List log stream",
	Long:  `Show all log stream avaliable on cloudwath region account`,
	Run: func(cmd *cobra.Command, args []string) {
		gname, err := cmd.Flags().GetString("gname")
		s, _ := cmd.Flags().GetInt64("sec")

		if err != nil {
			log.Fatalln("Error! Need log group name")
		}
		tools.Stream(gname, s)
	},
}

func init() {
	logCmd.AddCommand(streamCmd)
}

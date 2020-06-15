package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/thiago-scherrer/ocx/internal/tools"
)

// tailCmd represents the tail command
var tailCmd = &cobra.Command{
	Use:   "tail --gname LOG_GROUP_NAME --sname LOG_STREAM_NAME --lines 10",
	Short: "Subcommand to follow the update of the logs in real time",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		gname, err := cmd.Flags().GetString("gname")
		if err != nil {
			log.Fatalln("Error! Need Log Group name")
		}

		sname, err := cmd.Flags().GetString("sname")
		if err != nil {
			log.Fatalln("Error! Need Stream name")
		}

		lines, _ := cmd.Flags().GetInt64("lines")

		c := tools.Client()
		tools.Tail(c, gname, sname, lines)
	},
}

func init() {
	logCmd.AddCommand(tailCmd)
}

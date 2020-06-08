package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ocx",
	Short: "Another cli X",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().String("gname", "", "Group Name")

	rootCmd.PersistentFlags().String("sname", "", "Stream Name")

	rootCmd.PersistentFlags().Int64("lines", 1, "Lines")

	rootCmd.PersistentFlags().Int64("sec", 60, "Seconds")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}

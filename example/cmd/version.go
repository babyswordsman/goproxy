package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ouqiang/goproxy/example/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(version.Format())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

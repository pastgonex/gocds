package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &(cobra.Command{
	Use:   "version",
	Short: "Print the version number of Cds",
	Run: func(cmd *cobra.Command, args []string) {
		println("cds version is 1.0.0")
	},
})

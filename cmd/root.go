package cmd

import "github.com/spf13/cobra"

var RootCmd = &(cobra.Command{
	Use: "cds",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
})

package cmd

import "github.com/spf13/cobra"

// swagger to restful Use:   "s2r",api

func init() {
	RootCmd.AddCommand(s2r)
}

var s2r = &(cobra.Command{
	Use:   "s2r",
	Short: "swagger to restful api",
	Long:  `swagger to restful api`,
	Run: func(cmd *cobra.Command, args []string) {

	},
})

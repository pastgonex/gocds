package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func init() {
	RootCmd.AddCommand(s2r)
}

func swagger2Restful(swaggerFilePath string) (err error) {
	command := exec.Command("swagger", "serve", "-F=swagger", swaggerFilePath)
	if err = command.Run(); err != nil {
		fmt.Println(err.Error())
	}
	return err
}

var s2r = &(cobra.Command{
	Use:   "s2r",
	Short: "swagger to restful api",
	Long:  `swagger to restful api`,

	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case len(args) == 0:
			fmt.Println("please input swagger file path...")
		case len(args) == 1:
			if err := swagger2Restful(args[0]); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		default:
			fmt.Println("too many args...")
		}
	},
	Example: "cds s2r ./swaggerFiles/swagger.json",
})

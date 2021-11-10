package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

func init() {
	RootCmd.AddCommand(dirtreeCmd)
}

func Exists(path string) bool {
	_, err := os.Stat(path) // os.Stat 获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//deepth:目录的深度，如果包含一级子目录，则深度为2，依此类推。
func tree(dirPath string, deepth int) (err error) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	// Print the special prefix for the topmost directory
	if deepth == 1 {
		fmt.Printf("!---%s\n", filepath.Base(dirPath))
	}

	//目录分割符
	pathSep := string(os.PathSeparator)

	for _, file := range dir {
		// if subDir, recursive.
		if file.IsDir() {
			fmt.Printf("|")
			// Print tab spacing based on depth
			for i := 0; i < deepth; i++ {
				fmt.Printf("    |")
			}
			fmt.Printf("---%s\n", file.Name())
			err := tree(dirPath+pathSep+file.Name(), deepth+1)
			if err != nil {
				return err
			}
			continue
		}

		// if regular file , print it
		fmt.Printf("|")
		for i := 0; i < deepth; i++ {
			fmt.Printf("    |")
		}
		fmt.Printf("---%s\n", file.Name())
	}
	return nil
}

var dirtreeCmd = &(cobra.Command{
	Use:   "dirtree",
	Short: "All file names under the path provided by the tree output.",
	Run: func(cmd *cobra.Command, args []string) {
		//for i, v := range args{
		//	fmt.Printf("%d %s\n", i, v)
		var dirPath string
		switch {
		case len(args) == 0:
			dirPath = "."
		case len(args) == 1:
			dirPath = args[0]
		default:
			fmt.Println("Too many arguments.")
			os.Exit(1)
		}
		if Exists(dirPath) {
			err := tree(dirPath, 1)
			if err != nil {
				return
			}
		} else {
			fmt.Println("The path does not exist.")
			os.Exit(1)
		}
	},
	Example: "cds dirtree ./cmd",
})

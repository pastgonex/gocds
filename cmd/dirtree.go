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

//实现linux tree命令
//deepth:目录的深度，如果包含一级子目录，则深度为2，依此类推。
func tree(dirPath string, deepth int) (err error) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	//打印最上层目录的特殊前缀
	if deepth == 1 {
		fmt.Printf("!---%s\n", filepath.Base(dirPath))
	}

	//目录分割符
	pathSep := string(os.PathSeparator)

	for _, file := range dir {
		//如果是子目录，则递归调用tree命令，且deepth+1
		if file.IsDir() {
			fmt.Printf("|")
			//根据深度打印制表符间隔
			for i := 0; i < deepth; i++ {
				fmt.Printf("    |")
			}
			fmt.Printf("---%s\n", file.Name())
			tree(dirPath+pathSep+file.Name(), deepth+1)
			continue
		}

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
			tree(dirPath, 1)
		} else {
			fmt.Println("The path does not exist.")
			os.Exit(1)
		}
	},
})

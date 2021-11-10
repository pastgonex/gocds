package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	var dir string = "."
	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	}

	tree(dir, 1)
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

	//目录分割符，linux为\，windows为/
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

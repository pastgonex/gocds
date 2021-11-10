package main

import (
	"fmt"
	"gocds/cmd"
	"os"
)

type User struct {
	Id int `orm:"Id" json:"Id"`
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

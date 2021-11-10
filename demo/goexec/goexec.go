package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
)

func main() {
	command := exec.Command("ping", "www.baidu.com")
	outinfo := bytes.Buffer{}
	command.Stdout = &outinfo
	err := command.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
	if err = command.Wait(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(command.ProcessState.Pid())
		fmt.Println(command.ProcessState.Sys().(syscall.WaitStatus).ExitStatus())
		fmt.Println(outinfo.String())
	}
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	args := []string{""}
	binary, lookForErr := exec.LookPath("notepad.exe")
	if lookForErr != nil {
		panic("the binary not found." + lookForErr.Error())
	} else {
		fmt.Println(binary)
	}
	env := os.Environ()
	err := syscall.Exec(binary, args, env)
	if err != nil {
		panic(err)
	}
}

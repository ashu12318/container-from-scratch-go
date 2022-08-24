package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// go run main.go run <cmd> <params>
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("Unknown argument")
	}
}

func run() {
	//Restart from 10min
	fmt.Printf("Running %v\n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	cmd.Run()
}

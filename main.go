package main

import (
	"github.com/ChristianSchleifer/mremoteng/cmd"
/*	"os"
	"os/exec"
	"fmt"
	"syscall"*/
)	

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

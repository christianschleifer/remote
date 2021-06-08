package main

import (
	"github.com/ChristianSchleifer/mremoteng/cmd"
)	

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

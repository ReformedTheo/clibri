package main

import (
	"fmt"
	"os"

	"github.com/ReformedTheo/clibri/config"
)

const MAX_ARGS int = 5
const help string = "\n Usage: go run main.go help to see avaliable commands"

func main() {

	args := os.Args[1:]
	argCount := len(args)
	if argCount > MAX_ARGS {
		fmt.Println("Too many arguments!", help)
		os.Exit(1)
	} else if argCount <= 0 {
		fmt.Println("No argument found.", help)
		os.Exit(1)
	}

	switch args[0] {
	case "config":
		config.ConfigParser(args[2:])
	}

}

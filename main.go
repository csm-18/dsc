package main

import (
	"fmt"
	"os"
)

const VERSION = "v1.0.0"

func main() {
	var command string
	var args = os.Args

	if len(args) == 1 {
		command = "about"
	} else if len(args) == 2 {
		command = args[1]
	} else {
		fmt.Println("dsc: invalid command!")
		fmt.Println("On how to use dsc, run:\n dsc help")
		os.Exit(0)
	}

	user_prompt(command)
}

func user_prompt(command string) {
	println("user entered:", command)
}

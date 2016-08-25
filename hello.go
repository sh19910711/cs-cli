package main

import (
	"fmt"
)

var cmdHello = &Command{
	Dev:   true,
	Run:   runHello,
	Usage: "hello",
}

func runHello(cmd *Command, args []string) {
	fmt.Println("Hello", args)
}

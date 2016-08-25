package main

import (
	"fmt"
)

var cmdHello = &Command{
	Run:       runHello,
	UsageLine: "hello",
	Dev:       true,
}

func runHello(cmd *Command, args []string) {
	fmt.Println("Hello", args)
}

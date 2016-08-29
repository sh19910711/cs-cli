// play ground to test main.go
package main

import (
	"fmt"
)

var cmdHello = &Command{
	Dev:   true,
	Run:   runHello,
	Usage: "hello",
}

var helloHi bool

func init() {
	cmdHello.Flag.BoolVar(&helloHi, "hi", false, "")
}

func runHello(cmd *Command, args []string) error {
	if helloHi {
		fmt.Println("Hi", args)
	} else {
		fmt.Println("Hello", args)
	}

	return nil
}

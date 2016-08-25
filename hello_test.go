package main

import (
	"os"
)

// $ codestand hello world
func ExampleHelloWorld() {
	os.Args = []string{"cli", "hello", "world"}
	run()

	// Output:
	// Hello [world]
}

// $ codestand hello --hi world
func ExampleHelloHi() {
	os.Args = []string{"cli", "hello", "--hi", "world"}
	run()

	// Output:
	// Hi [world]
}

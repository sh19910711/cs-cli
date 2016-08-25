package main

import (
	"os"
)

// $ codestand hello
func ExampleHello() {
	os.Args = []string{"cli", "hello", "world"}
	run()

	// Output:
	// Hello [world]
}

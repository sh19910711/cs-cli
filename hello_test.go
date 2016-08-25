package main

import (
	"os"
)

// $ codestand hello
func ExampleHello() {
	os.Args = []string{"cli", "hello", "world"}
	main()

	// Output:
	// Hello [world]
}

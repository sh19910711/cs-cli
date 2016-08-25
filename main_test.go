package main

import (
	"os"
)

// $ codestand version
func ExampleVersion() {
	Version = "hello"
	os.Args = []string{"cli", "version"}
	main()

	// Output:
	// codestand/cli version hello
}

// $ codestand hello
func ExampleHello() {
	os.Args = []string{"cli", "hello", "world"}
	main()

	// Output:
	// Hello [world]
}

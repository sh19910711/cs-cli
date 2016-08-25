package main

import (
	"os"
)

// $ codestand version
func ExampleVersion() {
	Version = "hello"
	os.Args = []string{"cli", "version"}
	run()

	// Output:
	// codestand/cli version hello
}

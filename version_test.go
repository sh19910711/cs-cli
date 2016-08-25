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

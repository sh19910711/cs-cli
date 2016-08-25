package main

import (
	"os"
)

func ExampleVersion() {
	Version = "hello"
	os.Args = []string{"codestand", "version"}
	main()

	// Output:
	// codestand/cli version hello
}

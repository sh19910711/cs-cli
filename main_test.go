package main

import (
	"os"
)

func ExampleVersion() {
	Version = "hello"
	os.Args = []string{"cli", "version"}
	main()

	// Output:
	// codestand/cli version hello
}

func ExampleHello() {
	os.Args = []string{"cli", "hello"}
	main()

	// Output:
	// Hello []
}

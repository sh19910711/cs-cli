package main

import (
	"os"
)

func ExampleNew() {
	os.RemoveAll("./hello-world")
	os.Args = []string{"codestand", "new", "hello-world"}
	run()

	// Output:
	// create hello-world
	// create hello-world/application.yaml
	// create hello-world/main.cpp
}

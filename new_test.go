package main

import (
	"os"
)

func ExampleNew() {
	os.RemoveAll("./hello-world")
	Run([]string{"codestand", "new", "hello-world"})

	// Output:
	// create hello-world
	// create hello-world/application.yaml
	// create hello-world/main.cpp
}

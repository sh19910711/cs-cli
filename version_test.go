package main

func ExampleVersion() {
	Version = "hello"
	Run([]string{"codestand", "version"})

	// Output:
	// codestand/cli version hello
}

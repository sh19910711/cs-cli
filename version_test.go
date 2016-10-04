package main

func ExampleVersion() {
	Version = "hello"
	App.Run([]string{"codestand", "version"})

	// Output:
	// codestand/cli version hello
}

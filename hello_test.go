package main

func ExampleHello() {
	cmdHello.Run(cmdHello, []string{})
	// output:
	// Hello []
}

func ExampleHelloWorld() {
	cmdHello.Run(cmdHello, []string{"World"})
	// output:
	// Hello [World]
}

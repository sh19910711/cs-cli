//
// The "version" shows the version of the command.
//
package main

import (
	"fmt"
)

var cmdVersion = &Command{
	Run:   runVersion,
	Usage: "version",
	Short: "print version",
}

func runVersion(cmd *Command, args []string) {
	fmt.Printf("codestand/cli version %s\n", Version)
}

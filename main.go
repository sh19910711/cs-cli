package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// This variable can be overridden by `-ldflags "-X=main.Version=$VERSION"`.
var Version = "dev"

var commands = []*Command{
	cmdVersion,
	cmdHello,
	cmdNew,
}

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		renderErrorTemplate(mainUsageTemplate, commands)
		return 2
	}

	for _, c := range commands {
		if c.Name() == args[0] && c.Runnable() {
			c.Flag.Parse(args[1:])
			if err := c.Run(c, c.Flag.Args()); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				return 1
			}
			return 0
		}
	}

	fmt.Fprintln(os.Stderr, "unknown command")
	return 2
}

const mainUsageTemplate = `codestand/cli

Usage: codestand command [arguments]

The commands are:
{{range .}}{{if .Runnable}}{{if not .Dev}}  {{.Name}} - {{.Short}}
{{end}}{{end}}{{end}}
`

func ErrorMessage(msg string) error {
	return errors.New(msg)
}

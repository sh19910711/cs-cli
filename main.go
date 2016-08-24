package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

type Command struct {
	Run       func(cmd *Command, args []string)
	UsageLine string // The first word should stand for its command name.
	Short     string
}

// This variable can be overridden by `-ldflags "-X=main.Version=$VERSION"`.
var Version = "dev"

var commands = []*Command{
	cmdVersion,
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		usage()
		os.Exit(2)
	}

	for _, c := range commands {
		if c.name() == args[0] && c.runnable() {
			c.Run(c, []string{})
			return
		}
	}
}

var usageTemplate = `codestand/cli

Usage: codestand command [arguments]

The commands are:
`

func usage() {
	bw := bufio.NewWriter(os.Stderr)
	render(bw, usageTemplate, commands)
	bw.Flush()
}

func render(w io.Writer, text string, data interface{}) {
	t := template.New("tmpl")
}

func (c *Command) name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) runnable() bool {
	return c.Run != nil
}

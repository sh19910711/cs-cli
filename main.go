package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

type Command struct {
	Run       func(cmd *Command, args []string)
	UsageLine string // The first word should stand for its command name.
	Short     string
	Flag      flag.FlagSet
	Dev       bool // If true, the command does not appear in the command list of usage()
}

// This variable can be overridden by `-ldflags "-X=main.Version=$VERSION"`.
var Version = "dev"

var commands = []*Command{
	cmdVersion,
	cmdHello,
}

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		usage()
		return 2
	}

	for _, c := range commands {
		if c.Name() == args[0] && c.Runnable() {
			c.Flag.Parse(args[1:])
			c.Run(c, c.Flag.Args())
			return 0
		}
	}

	fmt.Fprintln(os.Stderr, "unknown command")
	return 2
}

var usageTemplate = `codestand/cli

Usage: codestand command [arguments]

The commands are:
{{range .}}{{if .Runnable}}{{if not .Dev}}  {{.Name}} - {{.Short}}
{{end}}{{end}}{{end}}
`

func usage() {
	render(os.Stderr, usageTemplate, commands)
}

func render(w io.Writer, text string, data interface{}) {
	t := template.New("tmpl")
	template.Must(t.Parse(text))
	err := t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) Runnable() bool {
	return c.Run != nil
}

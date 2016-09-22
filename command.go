package main

import (
	"flag"
	"strings"
)

type Command struct {
	// Each command should implement this function
	Run func(cmd *Command, args []string) error

	// Usage describes how to use the command
	// The first word stands for its command name.
	Usage string

	// Short is a short description in a line
	Short string

	// Flag handles command line options
	Flag flag.FlagSet

	// If Dev is true, the command does not appear in the command list of usage
	Dev bool
}

func (c *Command) Name() string {
	name := c.Usage
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) Runnable() bool {
	return c.Run != nil
}

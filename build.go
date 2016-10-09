package main

import (
	"os"
	"github.com/urfave/cli"
)

var buildCommand = cli.Command {
	Name: "build",
	Usage: "build the app",
	Action: doBuild,
}

func doBuild(c *cli.Context) error {

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	return RunCommand("docker", "run", "--delete", "-v", cwd + ":/app", "-t", "codestand/baseos")
}

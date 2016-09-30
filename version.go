package main

import (
	"fmt"
	"github.com/urfave/cli"
)

var versionCommand = cli.Command{
	Name: "version",
	Usage: "print version",
	Action: doVersion,
}

func doVersion(c *cli.Context) error {
	fmt.Printf("codestand/cli version %s\n", Version)
	return nil
}

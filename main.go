package main

import (
	"os"
	"fmt"
	"errors"
	"github.com/urfave/cli"
)

// This variable can be overridden by `-ldflags "-X=main.Version=$VERSION"`.
var Version = "dev"

var commands = []cli.Command{
	newCommand,
	registerCommand,
	configCommand,
	versionCommand,
	installCommand,
}

func Run(args []string) {
	app := cli.NewApp()
	app.Commands = commands
	app.Version = Version
	app.Usage = "Codestand CLI"

	err := app.Run(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	Run(os.Args)
}

func ErrorMessage(msg string) error {
	return errors.New(msg)
}

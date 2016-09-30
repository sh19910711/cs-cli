package main

import (
	"os"
	"errors"
	"github.com/urfave/cli"
)

// This variable can be overridden by `-ldflags "-X=main.Version=$VERSION"`.
var Version = "dev"

var commands = []cli.Command{
	newCommand,
	versionCommand,
	installCommand,
}

func main() {
	app := cli.NewApp()
	app.Commands = commands
	app.Version = Version
	app.Usage = "Codestand CLI"
	app.Run(os.Args)
}

func ErrorMessage(msg string) error {
	return errors.New(msg)
}

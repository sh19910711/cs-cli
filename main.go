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
	configCommand,
	versionCommand,
	installCommand,
}

func Run(args []string) {
	app := cli.NewApp()
	app.Commands = commands
	app.Version = Version
	app.Usage = "Codestand CLI"
	app.Run(args)
}

func main() {
	Run(os.Args)
}

func ErrorMessage(msg string) error {
	return errors.New(msg)
}

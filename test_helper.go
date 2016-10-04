package main

import (
	"github.com/urfave/cli"
)

var App *cli.App

func init() {
	App = cli.NewApp()
	App.Commands = commands
}

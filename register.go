package main

import (
	"fmt"
	"errors"
	"net/http"
	"github.com/urfave/cli"
	"github.com/Songmu/prompter"
)

var registerCommand = cli.Command {
	Name: "register",
	Usage: "Register an application",
	Action: doRegister,
}


func doRegister(c *cli.Context) error {

	name := c.Args().Get(0)
	if name == "" {
		name = prompter.Prompt("Application name", "")
	}

	status, _, err := InvokeAPI("POST", "/apps", map[string]string {
		"name": name,
	})

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New(fmt.Sprintf("server returned %v", status))
	}

	return nil
}

package main

import (
	"fmt"
	"errors"
	"net/http"
	"github.com/urfave/cli"
)

var deployImageCommand = cli.Command {
	Name: "deploy-image",
	Usage: "deploy the app image",
	Action: doDeployImage,
	Flags: []cli.Flag {
		cli.StringFlag {
			Name: "app-name",
			Value:"",
			Usage: "The application name.",
		},
		cli.StringFlag {
			Name: "image",
			Value:"",
			Usage: "The image file.",
		},
	},
}

func doDeployImage(c *cli.Context) error {
	appName := GetArgumentOrPrompt(c, "app-name", "Application Name", "")
	image := GetArgumentOrPrompt(c, "image", "The image file", "")

	status, _, err := InvokeAPI("POST", "/apps/" + appName + "/deployments",
		nil, map[string]string { "image": image })

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New(fmt.Sprintf("server returned %v", status))
	}

	return nil
}

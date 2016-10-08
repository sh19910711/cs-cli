package main

import (
	"fmt"
	"errors"
	"net/http"
	"github.com/urfave/cli"
)

var addDeviceCommand = cli.Command {
	Name: "add-device",
	Usage: "add a device to the app",
	Action: doAddDevice,
	Flags: []cli.Flag {
		cli.StringFlag {
			Name: "app-name",
			Value:"",
			Usage: "The application name.",
		},
		cli.StringFlag {
			Name: "device-name",
			Value:"",
			Usage: "The device name.",
		},
	},
}

func doAddDevice(c *cli.Context) error {

	appName := GetArgumentOrPrompt(c, "app-name", "Application Name", "")
	deviceName := GetArgumentOrPrompt(c, "device-name", "Device Name", "")

	status, _, err := InvokeAPI("POST", "/apps/" + appName + "/devices", map[string]string {
		"device": deviceName,
	}, "")

	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New(fmt.Sprintf("server returned %v", status))
	}

	return nil
}

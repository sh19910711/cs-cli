package main

import (
	"fmt"
	"errors"
	"encoding/json"
	"net/http"
	"github.com/urfave/cli"
)

var devicesCommand = cli.Command {
	Name: "devices",
	Usage: "List devices",
	Action: doDevices,
}

type ListDevicesResponse struct {
	Devices []struct {
		Name string
		Board string
		Status string
	}
}


func doDevices(c *cli.Context) error {

	status, body, err := InvokeAPI("GET", "/devices", nil, nil)
	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return errors.New(fmt.Sprintf("server returned %v", status))
	}

	var resp ListDevicesResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	for _, device := range resp.Devices {
		fmt.Printf("%v (%v): %v\n", device.Name, device.Board, device.Status)
	}

	return nil
}

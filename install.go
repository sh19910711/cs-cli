// Installs codestand OS to a device.
package main

import (
	"errors"
	"fmt"
	"path"
	"github.com/Songmu/prompter"
	"github.com/urfave/cli"
)

var installCommand = cli.Command {
	Name: "install",
	Usage: "install OS",
	Action: doInstall,
	Flags: []cli.Flag {
		cli.StringFlag {
			Name: "board",
			Value:"",
			Usage: "The board name. (esp8266)",
		},
		cli.StringFlag{
			Name: "serial",
			Value:"",
			Usage: "The serial port. (e.g. /dev/tty.USB0)",
		},
		cli.StringFlag{
			Name: "image",
			Value:"",
			Usage: "The firmware image file.",
		},
	},
}

func doInstall(c *cli.Context) error {
	board := c.String("board")
	serial := c.String("serial")
	image := c.String("image")

	// Prompts for those who does not know options.
	if board == "" {
		board = prompter.Prompt("Board Name", "esp8266")
	}

	if serial == "" {
		serial = prompter.Prompt("Serial port", DEFAULT_SERIAL_PORT)
	}

	// Validate inputs.
	var installer func(serial, image string) error
	var fwURL string

	if board == "esp8266" {
		installer = ESP8266Installer
		fwURL = ESP8266_FIRMWARE_URL
	} else {
		return errors.New(fmt.Sprintf("unknown board: '%v'", board))
	}

	if !FileExist(serial) {
		return errors.New(fmt.Sprintf("serial port '%v' does not exist", serial))
	}

	if image == "" {
		// The path to the image file is not specified; use the firmware image
		// on the GitHub.
		image = path.Join(GetConfigDir(), "cache", "firmware", board + ".img")
		if !FileExist(image) {
			PrepareConfigDir()
			DownloadFile(fwURL, image)
		}
	}

	return installer(serial, image)
}

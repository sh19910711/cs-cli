// Installs codestand OS to a device.
package main

import (
	"errors"
	"fmt"
	"path"
	"github.com/Songmu/prompter"
)

var cmdInstall = &Command{
	Dev:   false,
	Run:   runInstall,
	Usage: "install",
	Short: "install OS",
}


var board string
var serial string
var image string
func init() {
	cmdInstall.Flag.StringVar(&board,  "board",  "",  "The board name. (esp8266)")
	cmdInstall.Flag.StringVar(&serial, "serial", "",  "The serial port. (e.g. /dev/tty.USB0)")
	cmdInstall.Flag.StringVar(&image,  "image",  "",  "The firmware image file.")
}

func runInstall(cmd *Command, args []string) error {

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

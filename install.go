// Installs codestand OS to a device.
package main

import (
	"bytes"
	"errors"
	"fmt"
	"path"
	"os"
	"io/ioutil"
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
			Name: "device_name",
			Value:"",
			Usage: "The device name.",
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
		cli.StringFlag{
			Name: "wifi_ssid",
			Value:"",
			Usage: "The Wi-Fi SSID.",
		},
		cli.StringFlag{
			Name: "wifi_password",
			Value:"",
			Usage: "The Wi-Fi password.",
		},
	},
}

func replaceAndFillBytes(content []byte, old, new string) ([]byte, error) {
	old_bytes := []byte(old)
	new_bytes := []byte(new)
	old_len := len(old_bytes)
	new_len := len(new_bytes)

	if old_len - new_len < 0 {
		return nil, errors.New("replacement is too long")
	}

        replacement := make([]byte, old_len)
	copy(replacement, new_bytes)
	content = bytes.Replace(content, old_bytes, replacement, old_len)
	return content, nil
}

func getArgumentOrPrompt(c *cli.Context, flag, desc, def string) string {
	arg := c.String(flag)
	if arg == "" {
		arg = prompter.Prompt(desc, def)
	}

	return arg
}

func doInstall(c *cli.Context) error {
	board        := getArgumentOrPrompt(c, "board", "Board name", "esp8266")
	serial       := getArgumentOrPrompt(c, "serial", "Serial port", DEFAULT_SERIAL_PORT)
	image        := c.String("image")
	deviceName   := getArgumentOrPrompt(c, "device_name", "Device name", "")
	wifiSSID     := getArgumentOrPrompt(c, "wifi_ssid", "Wi-Fi SSID", "")
	wifiPassword := getArgumentOrPrompt(c, "wifi_password", "Wi-Fi password", "")

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
			if err := DownloadFile(fwURL, image); err != nil {
				return err
			}
		}
	}

	// Copy the original image to a temporary file.
	tmp_image, err := ioutil.TempFile("", "codestand")
	if err != nil {
		return err
	}

	defer os.Remove(tmp_image.Name())

	content, err := ioutil.ReadFile(image)
	if err != nil {
		return err
	}

	// Replace REPLACE_ME.
	content, err = replaceAndFillBytes(content, "__DEVICE_NAME__REPLACE_ME__", deviceName)
	content, err = replaceAndFillBytes(content, "__WIFI_SSID__REPLACE_ME__", wifiSSID)
	content, err = replaceAndFillBytes(content, "__WIFI_PASSWORD__REPLACE_ME__", wifiPassword)

	if _, err := tmp_image.Write(content); err != nil {
		return err
	}
	if err := tmp_image.Close(); err != nil {
		return err
	}

	return installer(serial, tmp_image.Name())
}

// Installs codestand OS to a device.
package main

import (
	"bytes"
	"errors"
	"fmt"
	"path"
	"os"
	"io/ioutil"
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
			Name: "device-name",
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
			Name: "wifi-ssid",
			Value:"",
			Usage: "The Wi-Fi SSID.",
		},
		cli.StringFlag{
			Name: "wifi-password",
			Value:"",
			Usage: "The Wi-Fi password.",
		},
	},
}

func replaceAndFillBytes(content []byte, old, new string) ([]byte, error) {
	oldBytes := []byte(old)
	newBytes := []byte(new)
	oldLen := len(oldBytes)
	newLen := len(newBytes)

	if oldLen - newLen < 0 {
		return nil, errors.New("replacement is too long")
	}

        replacement := make([]byte, oldLen)
	copy(replacement, newBytes)
	content = bytes.Replace(content, oldBytes, replacement, oldLen)
	return content, nil
}

func doInstall(c *cli.Context) error {
	board        := GetArgumentOrPrompt(c, "board", "Board name", "esp8266")
	serial       := GetArgumentOrPrompt(c, "serial", "Serial port", DEFAULT_SERIAL_PORT)
	image        := c.String("image")
	deviceName   := GetArgumentOrPrompt(c, "device-name", "Device name", "")
	wifiSSID     := GetArgumentOrPrompt(c, "wifi-ssid", "Wi-Fi SSID", "")
	wifiPassword := GetArgumentOrPromptPassword(c, "wifi-password", "Wi-Fi password")

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
	tmpImage, err := ioutil.TempFile("", "codestand")
	if err != nil {
		return err
	}

	defer os.Remove(tmpImage.Name())

	content, err := ioutil.ReadFile(image)
	if err != nil {
		return err
	}

	// Replace REPLACE_ME.
	content, err = replaceAndFillBytes(content, "__DEVICE_NAME__REPLACE_ME__", deviceName)
	content, err = replaceAndFillBytes(content, "__WIFI_SSID__REPLACE_ME__", wifiSSID)
	content, err = replaceAndFillBytes(content, "__WIFI_PASSWORD__REPLACE_ME__", wifiPassword)

	if _, err := tmpImage.Write(content); err != nil {
		return err
	}
	if err := tmpImage.Close(); err != nil {
		return err
	}

	return installer(serial, tmpImage.Name())
}

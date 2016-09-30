package main

import (
	"os"
	"os/exec"
)

const (
	ESP8266_FIRMWARE_URL = "https://github.com/codestand/esp8266-firmware/releases/download/v0.0.0/firmware.bin"
)

func ESP8266Installer(serial, image string) error {

	cmd := exec.Command("esptool", "-v", "-cd", "ck", "-cb", "115200",
		"-cp", serial, "-ca", "0x00000", "-cf", image)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

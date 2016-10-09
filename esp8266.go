package main

const (
	ESP8266_FIRMWARE_URL = "https://github.com/codestand/esp8266-firmware/releases/download/v0.0.1/firmware.bin"
)

func ESP8266Installer(serial, image string) error {

	return RunCommand("esptool", "-v", "-cd", "ck", "-cb", "115200",
		"-cp", serial, "-ca", "0x00000", "-cf", image)
}

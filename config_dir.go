package main

import (
	"os"
	"os/user"
	"path"
)

func GetConfigDir() string {
	u, _ := user.Current()
	return path.Join(u.HomeDir, ".codestand.d")
}

func PrepareConfigDir() {
	_ = os.Mkdir(GetConfigDir(), 0755)
	_ = os.Mkdir(path.Join(GetConfigDir(), "cache"), 0755)
	_ = os.Mkdir(path.Join(GetConfigDir(), "cache", "firmware"), 0755)
}

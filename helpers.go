package main
import (
	"os"
	"os/exec"
	"io/ioutil"
	"github.com/parnurzeal/gorequest"
	"github.com/urfave/cli"
	"github.com/Songmu/prompter"
)

func RunCommand(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

func GetArgumentOrPromptPassword(c *cli.Context, flag, desc string) string {
	arg := c.String(flag)
	if arg == "" {
		arg = prompter.Password(desc)
	}

	return arg
}

func GetArgumentOrPrompt(c *cli.Context, flag, desc, def string) string {
	arg := c.String(flag)
	if arg == "" {
		arg = prompter.Prompt(desc, def)
	}

	return arg
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func DownloadFile(url, path string) error {
	_, body, errs := gorequest.New().Get(url).EndBytes()
	ioutil.WriteFile(path, body, 0644)
	return errs[0]
}

package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

var newCommand = cli.Command {
	Name: "new",
	Usage: "create a new project directory",
	Action: doNew,
}

var applicationYamlTmpl = `api: 1
name: {{.AppName}}
lang: c++

sources:
  - main.cpp
`

var mainCppTmpl = `#include <app.h>

void setup() {

    // Add your cool code here!
}
`

var appName string

func createFile(path string, tmpl string, data interface{}) error {
	r, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := render(r, tmpl, data); err != nil {
		return err
	}
	return nil
}

func doNew(c *cli.Context) error {
	appName = c.Args().Get(0) // TODO: validate appName

	if err := os.Mkdir(appName, 0755); err != nil {
		return err
	}
	fmt.Printf("create %v\n", appName)
	if err := os.Chdir(appName); err != nil {
		return err
	}

	applicationYamlData := struct {
		AppName string
	}{
		appName,
	}
	if err := createFile("application.yaml", applicationYamlTmpl, applicationYamlData); err != nil {
		return err
	}
	fmt.Printf("create %s/application.yaml\n", appName)

	mainCppData := struct {
	}{}
	if err := createFile("main.cpp", mainCppTmpl, mainCppData); err != nil {
		return err
	}
	fmt.Printf("create %s/main.cpp\n", appName)

	return nil
}

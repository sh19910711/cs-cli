package main

import (
	"os"
)

var cmdNew = &Command{
	Run:   runNew,
	Usage: "new",
	Short: "create a new project directory",
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

const newUsageTemplate = `Usage: codestand new <app-name> [options]

Options:
  --template  USER/REPO    The template repository on GitHub.

`

func runNew(cmd *Command, args []string) error {
	if len(args) != 1 {
		renderErrorTemplate(newUsageTemplate, nil)
		return ErrorMessage("arguments error")
	}

	appName = args[0] // TODO: validate appName

	if err := os.Mkdir(appName, 0755); err != nil {
		return err
	}
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

	mainCppData := struct {
	}{}
	if err := createFile("main.cpp", mainCppTmpl, mainCppData); err != nil {
		return err
	}

	return nil
}

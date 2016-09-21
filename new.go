package main

import (
	"os"
	"bytes"
	"flag"
	"io/ioutil"
	"text/template"
)

var cmdNew = &Command{
	Run:   runNew,
	Usage: "new",
	Short: "create a new project directory",
}

var applicationYamlTmpl = `
api: 1
name: {{.AppName}}
lang: c++

sources:
  - main.cpp
`

var mainCppTmpl = `
#include <app.h>

void setup() {

    // Add your cool code here!
}
`

var appName string

func init() {
	flag.Parse()
	args := flag.Args()
	appName = args[1]

	// TODO: validate appName
}


func createFile(path string, tmpl string, data interface{}) {
	t, err := template.New("application yaml").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	var s bytes.Buffer
	err = t.Execute(&s, data)
	if err != nil {
		panic(err)
	}
	
	ioutil.WriteFile(path, s.Bytes(), 0644)
}


func runNew(cmd *Command, args []string) error {


	// TODO: error check
	os.Mkdir(appName, 0755)
	os.Chdir(appName)

	applicationYamlData := struct {
		AppName string
	}{
		appName,
	}
	createFile("application.yaml", applicationYamlTmpl, applicationYamlData)

	mainCppData := struct {
	}{
	}
	createFile("main.cpp", mainCppTmpl, mainCppData)
	return nil
}

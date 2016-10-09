package main

import (
	"fmt"
	"errors"
	"net/http"
	"path/filepath"
	"github.com/urfave/cli"
	"github.com/jhoonb/archivex"
)

var deployCommand = cli.Command {
	Name: "deploy",
	Usage: "deploy the app",
	Action: doDeploy,
	Flags: []cli.Flag {
		cli.StringFlag {
			Name: "app-name",
			Value:"",
			Usage: "The application name.",
		},
	},
}

func doDeploy(c *cli.Context) error {
	appName := GetArgumentOrPrompt(c, "app-name", "Application Name", "")

	sourceFile := ".app.zip"
	zip := new(archivex.ZipFile)
	zip.Create(sourceFile)

	// TODO: exclude unnecessary files
	for _, pattern := range []string{ "*", ".*" } {
		files, err := filepath.Glob(pattern)
		if err != nil {
			return nil
		}

		for _, file := range files {
			zip.AddFile(file)
		}
	}

	zip.Close()

	status, _, err := InvokeAPI("POST", "/apps/" + appName + "/builds",
		nil, map[string]string { "source_file": sourceFile })

	if err != nil {
		return err
	}

	if status != http.StatusAccepted {
		return errors.New(fmt.Sprintf("server returned %v", status))
	}

	return nil
}

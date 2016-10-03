package main

import (
	"io/ioutil"
	"fmt"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"github.com/parnurzeal/gorequest"
)

var configCommand = cli.Command {
	Name: "config",
	Usage: "create a build config",
	Action: doConfig,
}

type Config struct {
	Desc string
	Type string
	Default interface{} `yaml:"-"`
	Value interface{}
	Driver string `yaml:",omitempty"`
	Interface string `yaml:",omitempty"`
	Gpio struct {
		Pin int
	} `yaml:",omitempty"`
}

type AppYaml struct {
	Name string
	Api int
	Lang string
	Sources []string
	Cpp struct {
		Class_name string
	}
	Config map[string] Config
}

type DriverYaml struct {
	Name string
	Api int
	Lang string
	Sources []string
	Cpp struct {
		Class_name string
	}
	Config map[string] Config
}

func downloadDriverYaml(driver_name string) (error, *DriverYaml) {
	driverYaml := DriverYaml{}

	url := fmt.Sprintf("https://raw.githubusercontent.com/codestand/baseos/master/drivers/%s/driver.yaml",
		driver_name)
	_, body, errs := gorequest.New().Get(url).EndBytes()
	if errs != nil {
		return errs[0], nil
	}

	err := yaml.Unmarshal(body, &driverYaml)
	if err != nil {
		return err, nil
	}

	return nil, &driverYaml
}

func doConfig(c *cli.Context) error {

	// Predefined config
	configs := make(map[string] Config)
	configs["BAORD"] = Config {
		Type: "string",
		Default: "esp8266",
	}

	// Config in application.yaml
	filedata, err := ioutil.ReadFile("application.yaml")
	if err != nil {
		return err
	}

	appYaml := AppYaml{}
	err = yaml.Unmarshal(filedata, &appYaml)
	if err != nil {
		return err
	}

	for k, v := range appYaml.Config {
		configs[k] = v
	}

	for k, _ := range configs {
		config := configs[k]
		config.Value = config.Default
		config.Default = nil
		configs[k] = config
	}

	// Create .config.yaml
	data, err := yaml.Marshal(&configs)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(".config.yaml", data, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Edit .config.yaml!")
	return nil
}

package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jonog/redalert/core"
	"github.com/jonog/redalert/notifiers"
)

type Config struct {
	Checks        []core.CheckConfig `json:"checks"`
	Notifications []notifiers.Config `json:"notifications"`
}

func ReadConfigFile() (*Config, error) {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(file, &config)
	return &config, err
}

func ConfigureStdErr(s *core.Service) {
	s.Alerts["stderr"] = notifiers.NewStandardError()
}

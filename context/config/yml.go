package config

import (
	"github.com/rwirdemann/marketsync/ports/out"
	"gopkg.in/yaml.v3"
	"os"
)

type Yml struct{}

func (y Yml) ReadConfig() (out.Config, error) {
	data, err := os.ReadFile("marketsync.yaml")
	if err != nil {
		return out.Config{}, err
	}

	var c out.Config
	if err := yaml.Unmarshal(data, &c); err != nil {
		return out.Config{}, err
	}

	return c, nil
}

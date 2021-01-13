package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Pipeline struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type Config struct {
	Pipelines []Pipeline `yaml:"pipelines"`
}

var Conf Config

func init() {
	content, err := ioutil.ReadFile("engine.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, &Conf)
	if err != nil {
		panic(err)
	}
}

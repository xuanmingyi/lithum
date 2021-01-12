package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Pipeline string `yaml:"pipeline"`
}

type Pipeline struct {
	Name          string `yaml:"name"`
	Path          string `yaml:"path"`
	InputContent  string
	FilterContent string
	OutputContent string
}

func (p *Pipeline) Load() (err error) {
	path := fmt.Sprintf("%s/input.lua", p.Path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
		return err
	}
	p.InputContent = string(content)

	path = fmt.Sprintf("%s/filter.lua", p.Path)
	content, err = ioutil.ReadFile(path)
	if err != nil {
		panic(err)
		return err
	}
	p.FilterContent = string(content)

	path = fmt.Sprintf("%s/output.lua", p.Path)
	content, err = ioutil.ReadFile(path)
	if err != nil {
		panic(err)
		return err
	}
	p.OutputContent = string(content)

	return nil
}

var Conf *Config
var Pipelines []Pipeline

func LoadConfig(path string) {
	Conf = new(Config)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(content, Conf)
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	content, err = ioutil.ReadFile(Conf.Pipeline)
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(content, &Pipelines)
	if err != nil {
		panic(err)
		os.Exit(1)
	}

}

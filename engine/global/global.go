package global

import (
	"context"
	inputexec "engine/input/exec"
	inputhttp "engine/input/http"
	outputredis "engine/output/redis"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type pipeConfig struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type global struct {
	// 配置文件
	PipeConfs     []pipeConfig `yaml:"pipelines"`
	EnableMetrics bool         `yaml:"enable_metrics"`

	Pipelines []Pipeline
	GlobalCtx context.Context

	// modules
	InputModules  map[string]InputInitHandler
	OutputModules map[string]OutputInitHandler
}

var Global global

func init() {
	Global.Init()
}

func (g *global) Init() {
	content, err := ioutil.ReadFile("engine.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, g)
	if err != nil {
		panic(err)
	}
}

func RegistModules() {
	// Input
	RegistInputHandler(inputexec.ModuleName, inputexec.InitHandler)
	RegistInputHandler(inputhttp.ModuleName, inputhttp.InitHandler)

	// Output
	RegistOutputHandler(outputredis.ModuleName, outputredis.InitHandler)
}

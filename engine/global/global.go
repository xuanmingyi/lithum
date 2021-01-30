package global

import (
	"context"
	"gopkg.in/yaml.v2"
	"io/ioutil"

	inputcron "engine/input/cron"
	inputhttp "engine/input/http"
	outputmysql "engine/output/mysql"
	outputsqlite "engine/output/sqlite"
)

type pipeConfig struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type global struct {
	// 配置文件
	PipeConfs     []pipeConfig `yaml:"pipelines"`
	EnableMetrics bool         `yaml:"enable_metrics"`

	Pipelines []Pipeline      `yaml:"-"`
	GlobalCtx context.Context `yaml:"-"`

	// modules
	InputModules  map[string]InputInitHandler  `yaml:"-"`
	OutputModules map[string]OutputInitHandler `yaml:"-"`
}

var Global global

func init() {
	RegistModules()

	Global.Init()
}

func (g *global) Init() {
	content, err := ioutil.ReadFile("pipelines/engine.yaml")
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
	RegistInputHandler(inputhttp.ModuleName, inputhttp.InitHandler)
	RegistInputHandler(inputcron.ModuleName, inputcron.InitHandler)

	// Output
	RegistOutputHandler(outputmysql.ModuleName, outputmysql.InitHandler)
	RegistOutputHandler(outputsqlite.ModuleName, outputsqlite.InitHandler)
}

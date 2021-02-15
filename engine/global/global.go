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

// 全局结构
type global struct {
	// 配置文件
	PipelineConfs []struct {
		Name string `yaml:"name"`
		Path string `yaml:"path"`
	} `yaml:"pipelines"`

	Database struct {
		Driver string `yaml:"driver"`
		DSN    string `yaml:"dsn""`
	} `yaml:"database"`

	LibPath string `yaml:"lib_path"`

	Pipelines []*Pipeline `yaml:"-"`

	GlobalCtx context.Context    `yaml:"-"`
	Cancel    context.CancelFunc `yaml:"-"`

	// modules
	InputModules  map[string]InputInitHandler  `yaml:"-"`
	OutputModules map[string]OutputInitHandler `yaml:"-"`
}

// 全局变量
var Global global

func init() {
	// 注册所有模块
	RegistModules()

	// 读取配置文件
	content, err := ioutil.ReadFile("pipelines/engine.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, &Global)
	if err != nil {
		panic(err)
	}

	// 初始化cancelCtx
	Global.GlobalCtx, Global.Cancel = context.WithCancel(context.Background())

	// 初始化 Pipelines
	for _, conf := range Global.PipelineConfs {
		_pipeline, err := LoadPipeline(conf.Name, conf.Path)
		if err != nil {
			panic(err)
		}
		Global.Pipelines = append(Global.Pipelines, _pipeline)
	}
}

func (g *global) StartAll() {
	for _, pipeline := range g.Pipelines {
		go pipeline.Start()
	}
}

func (g *global) WaitAll() {
	for _, pipeline := range Global.Pipelines {
		pipeline.Wait()
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

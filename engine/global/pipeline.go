package global

import (
	"context"
	"engine/models"
	"engine/pipeline"
	"io/ioutil"
	"path"
	"sync"

	"gopkg.in/yaml.v2"
)

type Pipeline struct {
	Name       string
	InputRaw   string
	FilterRaw  string
	OutputRaw  string
	Inputs     []pipeline.Input
	Filter     pipeline.Filter
	Outputs    []pipeline.Output
	InputChan  chan models.Event
	OutputChan chan models.Event
	wg         sync.WaitGroup
	Ctx        context.Context
}

func (p *Pipeline) Start() {
	for _, input := range p.Inputs {
		p.wg.Add(1)
		go input.Start(p.InputChan)
	}

	p.wg.Add(1)
	go p.Filter.Start(p.InputChan, p.OutputChan)

	for _, output := range p.Outputs {
		p.wg.Add(1)
		go output.Start(p.OutputChan)
	}
}

func (p *Pipeline) Wait() {
	p.wg.Wait()
}

func ReadFile(pipeConf *pipeConfig, name string) string {
	content, err := ioutil.ReadFile(path.Join(pipeConf.Path, name))
	if err != nil {
		panic(err)
	}
	return string(content)
}

func LoadPipeline(pipeConf *pipeConfig) (pipeline *Pipeline, err error) {

	pipeline = new(Pipeline)

	pipeline.Ctx = Global.GlobalCtx

	pipeline.InputChan = make(chan models.Event, 1000)
	pipeline.OutputChan = make(chan models.Event, 1000)

	pipeline.Name = pipeConf.Name

	// Read Files
	pipeline.InputRaw = ReadFile(pipeConf, "input.yaml")
	pipeline.FilterRaw = ReadFile(pipeConf, "filter.lua")
	pipeline.Filter.Code = pipeline.FilterRaw
	pipeline.OutputRaw = ReadFile(pipeConf, "output.yaml")

	// Load Input
	var inputConfigs []map[string]interface{}
	err = yaml.Unmarshal([]byte(pipeline.InputRaw), &inputConfigs)
	if err != nil {
		panic(err)
	}

	for _, inputConfig := range inputConfigs {
		handler := Global.InputModules[inputConfig["driver"].(string)]
		input, err := handler(pipeConf, inputConfig)
		if err != nil {
			panic(err)
		}
		pipeline.Inputs = append(pipeline.Inputs, input)
	}

	// Load Output
	var outputConfigs []map[string]interface{}
	err = yaml.Unmarshal([]byte(pipeline.OutputRaw), &outputConfigs)
	if err != nil {
		panic(err)
	}

	for _, outputConfig := range outputConfigs {
		handler := Global.OutputModules[outputConfig["driver"].(string)]
		output, err := handler(pipeConf, outputConfig)
		if err != nil {
			panic(err)
			pipeline.Outputs = append(pipeline.Outputs, output)
		}
	}

	return pipeline, nil
}

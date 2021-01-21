package global

import (
	"context"
	inputexec "engine/input/exec"
	inputhttp "engine/input/http"
	"engine/models"
	outputmysql "engine/output/mysql"
	outputredis "engine/output/redis"
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

func LoadPipeline(p *Pipeline, ctx context.Context) (pipeline *Pipeline, err error) {
	pipeline = new(Pipeline)

	pipeline.Ctx = ctx

	pipeline.InputChan = make(chan models.Event, 1000)
	pipeline.OutputChan = make(chan models.Event, 1000)

	pipeline.Name = p.Name

	content, err := ioutil.ReadFile(path.Join(p.Path, "input.yaml"))
	if err != nil {
		panic(err)
	}
	pipeline.InputRaw = string(content)

	content, err = ioutil.ReadFile(path.Join(p.Path, "filter.lua"))
	if err != nil {
		panic(err)
	}
	pipeline.FilterRaw = string(content)
	pipeline.Filter.Code = pipeline.FilterRaw

	content, err = ioutil.ReadFile(path.Join(p.Path, "output.yaml"))
	if err != nil {
		panic(err)
	}
	pipeline.OutputRaw = string(content)

	var inputConfigs []map[string]string
	err = yaml.Unmarshal([]byte(pipeline.InputRaw), &inputConfigs)
	if err != nil {
		panic(err)
	}

	for _, inputConfig := range inputConfigs {
		switch inputConfig["driver"] {
		case "http":
			input, _ := inputhttp.InitHandler(pipeline.Ctx, inputConfig)
			pipeline.Inputs = append(pipeline.Inputs, input)
		case "exec":
			input, _ := inputexec.InitHandler(pipeline.Ctx, inputConfig)
			pipeline.Inputs = append(pipeline.Inputs, input)
		}
	}

	var outputConfigs []map[string]interface{}
	err = yaml.Unmarshal([]byte(pipeline.OutputRaw), &outputConfigs)
	if err != nil {
		panic(err)
	}

	for _, outputConfig := range outputConfigs {
		switch outputConfig["driver"].(string) {
		case "mysql":
			output, _ := outputmysql.InitHandler(pipeline.Ctx, outputConfig)
			pipeline.Outputs = append(pipeline.Outputs, output)
		case "redis":
			output, _ := outputredis.InitHandler(pipeline.Ctx, outputConfig)
			pipeline.Outputs = append(pipeline.Outputs, output)
		}
	}

	return pipeline, nil
}

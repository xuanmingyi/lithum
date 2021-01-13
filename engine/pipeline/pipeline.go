package pipeline

import (
	"context"
	"engine/config"
	inputexec "engine/input/exec"
	inputhttp "engine/input/http"
	"engine/models"
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
	Inputs     []Input
	Filter     Filter
	Outputs    []Output
	InputChan  chan models.Message
	OutputChan chan models.Message
	wg         sync.WaitGroup
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

func LoadPipeline(p *config.Pipeline) (pipeline *Pipeline, err error) {
	pipeline = new(Pipeline)

	pipeline.InputChan = make(chan models.Message, 1000)
	pipeline.OutputChan = make(chan models.Message, 1000)

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

	content, err = ioutil.ReadFile(path.Join(p.Path, "output.lua"))
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
			input, _ := inputhttp.InitHandler(context.Background(), inputConfig)
			pipeline.Inputs = append(pipeline.Inputs, input)
		case "exec":
			input, _ := inputexec.InitHandler(context.Background(), inputConfig)
			pipeline.Inputs = append(pipeline.Inputs, input)
		}
	}

	return pipeline, nil
}

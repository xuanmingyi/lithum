package pipeline

import (
	"engine/config"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"io/ioutil"
	"path"
	"sync"
)

type Message struct {
	L    *lua.LState
	Body string
}

type Pipeline struct {
	Name       string
	InputRaw   string
	FilterRaw  string
	OutputRaw  string
	Inputs     []Input
	Filter     Filter
	Outputs    []Output
	InputChan  chan Message
	OutputChan chan Message
	wg         sync.WaitGroup
}

func (p *Pipeline) Start() {
	for _, input := range p.Inputs {
		p.wg.Add(1)
		go input.Start(p.InputChan)
	}

	if p.Filter != nil {
		p.wg.Add(1)
		go p.Filter.Start(p.InputChan, p.OutputChan)
	}

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

	pipeline.Name = p.Name

	content, err := ioutil.ReadFile(path.Join(p.Path, "input.lua"))
	if err != nil {
		panic(err)
	}
	pipeline.InputRaw = string(content)

	content, err = ioutil.ReadFile(path.Join(p.Path, "filter.lua"))
	if err != nil {
		panic(err)
	}
	pipeline.FilterRaw = string(content)

	content, err = ioutil.ReadFile(path.Join(p.Path, "output.lua"))
	if err != nil {
		panic(err)
	}
	pipeline.OutputRaw = string(content)

	L := lua.NewState()
	defer L.Close()
	err = L.DoString(pipeline.InputRaw)
	if err != nil {
		panic(err)
	}
	inputResult := L.GetGlobal("input").(*lua.LTable)

	fmt.Println(inputResult.RawGetString("driver").String())

	//	inputResult.ForEach(func(value lua.LValue, value2 lua.LValue) {
	//		fmt.Println(value, value2)
	//	})

	return pipeline, nil
}

package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Task struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}


func HttpClient(url string) string{
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}


func RunPipeline(task Task, group *sync.WaitGroup) {
	defer group.Done()

	inputPath := fmt.Sprintf("%s/input.lua", task.Path)
	filterPath := fmt.Sprintf("%s/filter.lua", task.Path)
	outputPath := fmt.Sprintf("%s/output.lua", task.Path)

	L := lua.NewState()
	defer L.Close()

	if err := L.DoFile(inputPath); err != nil {
		panic(err)
	}
	input := L.GetGlobal("input")

	driver := L.GetField(input, "driver").(lua.LString)

	if driver == "http" {
		ticker := time.NewTicker(time.Second * 30)
		for {
			select {
			case <-ticker.C:
					html := HttpClient(string(L.GetField(input, "url").(lua.LString)))
					L.SetGlobal("message", lua.LString(html))
					if err := L.DoFile(filterPath); err != nil {
						panic(err)
						return
					}
					if err := L.DoFile(outputPath); err != nil {
						panic(err)
						return
					}
					break
			}
		}
	}
}

func main() {
	var tasks *[]Task
	wg := sync.WaitGroup{}
	tasks = new([]Task)

	content, err := ioutil.ReadFile("pipelines.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, tasks)
	if err != nil {
		panic(err)
	}

	for _, v := range *tasks {
		wg.Add(1)
		go RunPipeline(v, &wg)
	}
	wg.Wait()


}
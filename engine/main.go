package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"

	mysql "github.com/tengattack/gluasql/mysql"
)

type Task struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}



func RunPipeline(task Task, group *sync.WaitGroup) {
	defer group.Done()

	path := fmt.Sprintf("%s/main.lua", task.Path)
	fmt.Println(path)

	L := lua.NewState()
	L.PreloadModule("mysql", mysql.Loader)
	defer L.Close()

	if err := L.DoString(`print("hello world")`); err != nil {
		panic(err)
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
package cmd

import (
	"engine/config"
	"flag"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"sync"
)

var configPath string

func RunPipeline(pipeline config.Pipeline, group *sync.WaitGroup) (err error) {
	defer group.Done()
	pipeline.Load()
	L := lua.NewState()
	defer L.Close()
	err = L.DoString(pipeline.InputContent)
	if err != nil {
		panic(err)
		return
	}
	input := L.GetGlobal("input")
	driverName := L.GetField(input, "driver").(lua.LString).String()

	// 启动Driver的Input

	inputMaps[driverName].Start()
	fmt.Println(driver)

	return nil
}

func MainRun() {
	const (
		defaultConfig = "engine.yaml"
	)

	flag.StringVar(&configPath, "c", defaultConfig, "配置文件")

	flag.Parse()

	config.LoadConfig(configPath)

	var wg sync.WaitGroup

	for _, pipeline := range config.Pipelines {
		wg.Add(1)
		go RunPipeline(pipeline, &wg)
	}

	wg.Wait()
}

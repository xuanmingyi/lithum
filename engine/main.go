package main

import (
	"engine/config"
	"engine/pipeline"
)

func main() {
	var pipelines []*pipeline.Pipeline

	for _, p := range config.Conf.Pipelines {
		_pipeline, err := pipeline.LoadPipeline(&p)
		if err != nil {
			panic(err)
		}
		pipelines = append(pipelines, _pipeline)
	}
	for _, p := range pipelines {
		go p.Start()
	}

	for _, p := range pipelines {
		p.Wait()
	}

	for {
	}
}

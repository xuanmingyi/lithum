package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"engine/config"
	"engine/pipeline"
)

func waitSingals(ctx context.Context) error {
	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, os.Interrupt, os.Kill)

	select {
	case <-ctx.Done():
	case sig := <-osSignalChan:
		fmt.Println(sig)
	}
	return nil
}

func main() {
	var pipelines []*pipeline.Pipeline

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, p := range config.Conf.Pipelines {
		_pipeline, err := pipeline.LoadPipeline(&p, ctx)
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

	waitSingals(ctx)
}

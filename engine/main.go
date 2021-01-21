package main

import (
	"context"
	"engine/metric"
	"fmt"
	"os"
	"os/signal"

	"engine/global"
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
	var pipelines []*global.Pipeline

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, p := range global.Config.Pipelines {
		_pipeline, err := global.LoadPipeline(&p, ctx)
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

	if global.Config.EnableMetrics {
		go metric.Run()
	}

	waitSingals(ctx)
}

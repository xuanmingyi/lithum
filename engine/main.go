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

	global.Global.GlobalCtx = ctx

	for _, pipeConf := range global.Global.PipeConfs {
		pipeline, err := global.LoadPipeline(&pipeConf)
		if err != nil {
			panic(err)
		}
		pipelines = append(pipelines, pipeline)
	}

	for _, pipeline := range pipelines {
		go pipeline.Start()
	}

	if global.Global.EnableMetrics {
		go metric.Run()
	}

	for _, pipeline := range pipelines {
		pipeline.Wait()
	}

	waitSingals(ctx)
}

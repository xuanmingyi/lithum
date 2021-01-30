package inputcron

import (
	"context"
	"engine/models"
	"engine/pipeline"
	"time"
)

const ModuleName = "cron"

type InputCron struct {
	Ctx      context.Context
	Interval int
}

func InitHandler(ctx context.Context, values map[string]interface{}) (input pipeline.Input, err error) {
	inputCron := new(InputCron)
	inputCron.Ctx = ctx
	inputCron.Interval = values["interval"].(int)
	return inputCron, nil
}

func (t *InputCron) Start(model chan models.Event) {
	startChan := make(chan bool, 1)
	ticker := time.NewTicker(time.Duration(t.Interval) * time.Second)
	startChan <- true

	for {
		select {
		case <-t.Ctx.Done():
			return
		case <-startChan:
			model <- models.Event{}
		case <-ticker.C:
			model <- models.Event{}
		}
	}
}

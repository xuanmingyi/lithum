package inputexec

import (
	"context"
	"engine/models"
	"strconv"
	"time"
)

const ModuleName = "exec"

const ErrorTag = "engine_input_exec_error"

type InputExec struct {
	Ctx      context.Context
	Cmd      string
	Interval int
}

func InitHandler(ctx context.Context, values map[string]string) (input *InputExec, err error) {
	input = new(InputExec)
	input.Ctx = ctx
	input.Interval, err = strconv.Atoi(values["interval"])
	if err != nil {
		return nil, err
	}
	input.Cmd = values["cmd"]
	return input, nil
}

func (t *InputExec) Start(msg chan models.Message) {
	startChan := make(chan bool, 1)
	ticker := time.NewTicker(time.Duration(t.Interval) * time.Second)
	startChan <- true
	for {
		select {
		case <-startChan:
			msg <- t.Exec()
		case <-ticker.C:
			msg <- t.Exec()
		}
	}
}

func (t *InputExec) Exec() (message models.Message) {
	message.Body = "sssssssssssss"
	return message
}

package inputexec

import (
	"context"
	"engine/models"
	"engine/pipeline"
	"os/exec"
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

func InitHandler(ctx context.Context, values map[string]string) (input *pipeline.Input, err error) {
	inputExec := new(InputExec)
	inputExec.Ctx = ctx
	inputExec.Interval, err = strconv.Atoi(values["interval"])
	if err != nil {
		panic(err)
	}
	inputExec.Cmd = values["cmd"]
	return inputExec.(*pipeline.Input), nil
}

func (t *InputExec) Start(model chan models.Event) {
	startChan := make(chan bool, 1)
	ticker := time.NewTicker(time.Duration(t.Interval) * time.Second)
	startChan <- true

	for {
		select {
		case <-t.Ctx.Done():
			return
		case <-startChan:
			model <- t.Exec()
		case <-ticker.C:
			model <- t.Exec()
		}
	}
}

func (t *InputExec) Exec() (event models.Event) {
	cmd := exec.Command(t.Cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		event.Body = ""
		event.Tags = append(event.Tags, ErrorTag)
	} else {
		event.Body = string(out)
	}
	return event
}

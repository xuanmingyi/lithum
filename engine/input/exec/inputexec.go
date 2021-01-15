package inputexec

import (
	"context"
	"engine/models"
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

func InitHandler(ctx context.Context, values map[string]string) (input *InputExec, err error) {
	input = new(InputExec)
	input.Ctx = ctx
	input.Interval, err = strconv.Atoi(values["interval"])
	if err != nil {
		panic(err)
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
		case <-t.Ctx.Done():
			return
		case <-startChan:
			msg <- t.Exec()
		case <-ticker.C:
			msg <- t.Exec()
		}
	}
}

func (t *InputExec) Exec() (message models.Message) {
	cmd := exec.Command(t.Cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		message.Body = ""
		message.Tags = append(message.Tags, ErrorTag)
	} else {
		message.Body = string(out)
	}
	return message
}

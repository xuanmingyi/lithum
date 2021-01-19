package inputhttp

import (
	"context"
	"engine/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const ModuleName = "http"

const ErrorTag = "engine_input_http_error"

type InputHttp struct {
	Ctx      context.Context
	Interval int    `yaml:"interval"`
	URL      string `yaml:"url"`
}

func InitHandler(ctx context.Context, values map[string]string) (input *InputHttp, err error) {
	input = new(InputHttp)
	input.Ctx = ctx
	input.Interval, err = strconv.Atoi(values["interval"])
	if err != nil {
		return nil, err
	}
	input.URL = values["url"]
	return input, nil
}

func (t *InputHttp) Start(msg chan models.Event) {
	startChan := make(chan bool, 1)
	ticker := time.NewTicker(time.Duration(t.Interval) * time.Second)
	defer ticker.Stop()

	startChan <- true

	for {
		select {
		case <-startChan:
			msg <- t.Request()
		case <-ticker.C:
			msg <- t.Request()
		}
	}
}

func (t *InputHttp) Request() (event models.Event) {
	resp, err := http.Get(t.URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	event.Body = string(body)
	return event
}

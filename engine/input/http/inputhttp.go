package inputhttp

import (
	"context"
	"engine/models"
	"io/ioutil"
	"net/http"
	"time"
)

const ModuleName = "http"

const ErrorTag = "engine_input_http_error"

type InputHttp struct {
	Ctx      context.Context
	Interval int    `yaml:"interval"`
	URL      string `yaml:"url"`
}

func InitHandler(ctx context.Context, values map[string]interface{}) (i interface{ Start(chan models.Event) }, err error) {
	input := new(InputHttp)
	input.Ctx = ctx
	input.Interval = values["interval"].(int)
	input.URL = values["url"].(string)
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

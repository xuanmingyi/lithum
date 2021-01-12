package inputhttp

import (
	"context"
	"fmt"
	"time"
)

const ModuleName = "http"

const ErrorTag = "engine_input_http_error"

type InputConfig struct {
	Interval int `json:"interval"`
}

func DefaultInputConfig() InputConfig {
	return InputConfig{}
}

func InitHandler(ctx context.Context) {
}

func (t *InputConfig) Start(ctx context.Context) (err error) {
	startChan := make(chan bool, 1)
	ticker := time.NewTicker(time.Duration(t.Interval) * time.Second)
	defer ticker.Stop()

	startChan <- true

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-startChan:
			t.Request(ctx)
		case <-ticker.C:
			t.Request(ctx)
		}
	}
}

func (t *InputConfig) Request(ctx context.Context) {
	data, err := t.SendRequest()
	fmt.Println(data)
	fmt.Println(err)
}

func (t *InputConfig) SendRequest() (data []byte, err error) {
	return nil, nil
}

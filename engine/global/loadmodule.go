package global

import (
	"context"
	"engine/pipeline"
)

type InputInitHandler func(ctx context.Context, values map[string]string) (*pipeline.Input, error)
type OutputInitHandler func(ctx context.Context, values map[string]string) (*pipeline.Output, error)

func RegistInputHandler(name string, handler InputInitHandler) {
	Global.InputModules[name] = handler
}

func RegistOutputHandler(name string, handler OutputInitHandler) {
	Global.OutputModules[name] = handler
}

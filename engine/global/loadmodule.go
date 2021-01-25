package global

import (
	"context"
	"engine/log"
	"engine/pipeline"
	"github.com/sirupsen/logrus"
)

type InputInitHandler func(ctx context.Context, values map[string]interface{}) (pipeline.Input, error)
type OutputInitHandler func(ctx context.Context, values map[string]interface{}) (pipeline.Output, error)

func RegistInputHandler(name string, handler InputInitHandler) {
	if Global.InputModules == nil {
		Global.InputModules = make(map[string]InputInitHandler)
	}
	log.Log.WithFields(logrus.Fields{
		"module": name,
		"type":   "input",
	}).Info("注册模块")
	Global.InputModules[name] = handler
}

func RegistOutputHandler(name string, handler OutputInitHandler) {
	if Global.OutputModules == nil {
		Global.OutputModules = make(map[string]OutputInitHandler)
	}
	log.Log.WithFields(logrus.Fields{
		"module": name,
		"type":   "output",
	}).Info("注册模块")
	Global.OutputModules[name] = handler
}

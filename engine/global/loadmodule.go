package global

import (
	"context"
	"engine/log"
	"engine/models"
	"github.com/sirupsen/logrus"
)

type InputInitHandler func(ctx context.Context, values map[string]interface{}) (interface{ Start(chan models.Event) }, error)
type OutputInitHandler func(ctx context.Context, values map[string]interface{}) (interface{ Start(chan models.Event) }, error)

func RegistInputHandler(name string, handler InputInitHandler) {
	if Global.InputModules == nil {
		Global.InputModules = make(map[string]InputInitHandler)
	}
	Global.InputModules[name] = handler
	log.Log.WithFields(logrus.Fields{
		"module": name,
		"type":   "input",
	}).Info("注册输入模块")
}

func RegistOutputHandler(name string, handler OutputInitHandler) {
	if Global.OutputModules == nil {
		Global.OutputModules = make(map[string]OutputInitHandler)
	}
	Global.OutputModules[name] = handler
	log.Log.WithFields(logrus.Fields{
		"module": name,
		"type":   "output",
	}).Info("注册输出模块")
}

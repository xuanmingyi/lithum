package pipeline

import (
	"engine/models"

	lua "github.com/yuin/gopher-lua"
)

type Filter struct {
	Code string
}

func (f *Filter) Start(i chan models.Message, o chan models.Message) {
	var m models.Message

	for {
		m = <-i
		m.L = lua.NewState()
		m.L.SetGlobal("message", lua.LString(m.Body))
		err := m.L.DoString(f.Code)
		if err != nil {
			panic(err)
		}
		o <- m
	}
}

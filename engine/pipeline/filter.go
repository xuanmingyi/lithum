package pipeline

import (
	"engine/models"
	"fmt"
	libs "github.com/vadv/gopher-lua-libs"

	lua "github.com/yuin/gopher-lua"
)

type Filter struct {
	Code string
}

func (f *Filter) Start(input chan models.Message, output chan models.Message) {
	var msg models.Message

	for {
		msg = <-input
		if len(msg.Tags) == 0 {
			L := lua.NewState()
			libs.Preload(L)
			L.SetGlobal("message", lua.LString(msg.Body))
			err := L.DoString(f.Code)
			if err != nil {
				panic(err)
			}
			msg.Output = L.GetGlobal("output").(lua.LString).String()
			output <- msg
		} else {
			fmt.Println(msg)
		}
	}
}

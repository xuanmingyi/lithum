package models

import lua "github.com/yuin/gopher-lua"

type Message struct {
	L    *lua.LState
	Body string
}

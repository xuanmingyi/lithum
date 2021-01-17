package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

type Event struct {
	Values map[string]string
}

const luaEventTypeName = "event"

func registerEventType(L *lua.LState) {
	mt := L.NewTypeMetatable(luaEventTypeName)
	L.SetGlobal("event", mt)
	L.SetField(mt, "new", L.NewFunction(newEvent))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), eventMethods))
}

func newEvent(L *lua.LState) int {
	event := &Event{}
	ud := L.NewUserData()
	ud.Value = event
	L.SetMetatable(ud, L.GetTypeMetatable(luaEventTypeName))
	L.Push(ud)
	return 1
}

func checkEvent(L *lua.LState) *Event {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Event); ok {
		return v
	}
	return nil
}

var eventMethods = map[string]lua.LGFunction{
	"get": eventGet,
	"set": eventSet,
}

func eventGet(L *lua.LState) int {
	e := checkEvent(L)
	name := L.CheckString(2)
	L.Push(lua.LString(e.Values[name]))
	return 1
}

func eventSet(L *lua.LState) int {
	e := checkEvent(L)
	name := L.CheckString(2)
	value := L.CheckString(3)
	if e.Values == nil {
		e.Values = make(map[string]string, 100)
	}
	e.Values[name] = value
	return 0
}

func main() {
	L := lua.NewState()
	defer L.Close()
	registerEventType(L)
	err := L.DoString(`
a=event:new()
a:set("abc", "xuanmingyi")
print(a:get("abc"))
`)
	if err != nil {
		fmt.Println(err)
	}
}

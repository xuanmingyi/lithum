package models

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"strings"
)

type Event struct {
	Values map[string]interface{}
	Body   string
	Output string
	Tags   []string
}

const luaEventTypeName = "event"

func newEvent(L *lua.LState) {
	event := &Event{}

	mt := L.NewTypeMetatable(luaEventTypeName)
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), eventMethods))

	ud := L.NewUserData()
	ud.Value = event
	L.SetMetatable(ud, L.GetTypeMetatable(luaEventTypeName))

	L.SetGlobal("event", ud)
}

func checkEvent(L *lua.LState) (event *Event) {
	ud := L.CheckUserData(1)
	if event, ok := ud.Value.(*Event); ok {
		return event
	}
	return nil
}

func eventGet(L *lua.LState) int {
	event := checkEvent(L)
	fmt.Println(event)
	return 0
}

func splitName(name string) []string {
	var results []string
	var index int
	var count int = 0

	// 空字符串
	name = strings.TrimSpace(name)
	if name == "" {
		return nil
	}

	// 是否包含'[' ']'
	if !strings.Contains(name, "[") && !strings.Contains(name, "]") {
		results = append(results, name)
		return results
	}

	// 是否以'[' ']'开头结尾
	if !strings.HasPrefix(name, "[") || !strings.HasSuffix(name, "]") {
		return nil
	}

	// check '[' ']' 配对
	for i, v := range name {
		if v == '[' {
			count++
			index = i
		} else if v == ']' {
			count--
			results = append(results, name[index+1:i])
		}

		if (count >= 0) && (count <= 1) {
			continue
		} else {
			return nil
		}
	}
	return results
}

func eventSet(L *lua.LState) int {
	event := checkEvent(L)

	names := splitName(L.CheckString(2))

	if len(names) == 0 {
		panic("error")
		return 0
	}

	value := L.CheckAny(3)

	var pointer map[string]interface{}
	pointer = event.Values

	for i := 0; i < len(names)-1; i++ {
		if pointer[names[i]] == nil {
			pointer[names[i]] = make(map[string]interface{})
		}
		pointer = pointer[names[i]].(map[string]interface{})
	}

	switch value.(type) {
	case lua.LString:
		fmt.Println("LString")
	case lua.LNumber:
		fmt.Println("LNumber")
	case lua.LBool:
		fmt.Println("LBool")
	}
	fmt.Println(event, value)
	return 0
}

func eventDump(L *lua.LState) int {
	event := checkEvent(L)
	fmt.Println(event)
	return 0
}

var eventMethods = map[string]lua.LGFunction{
	"get":  eventGet,
	"set":  eventSet,
	"dump": eventDump,
}

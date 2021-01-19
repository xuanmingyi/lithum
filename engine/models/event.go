package models

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"strings"
)

type Event struct {
	Body   string
	Values map[string]interface{}
	Tags   []string
}

const luaEventTypeName = "event"

func NewEvent(L *lua.LState) (event *Event) {
	event = &Event{}

	mt := L.NewTypeMetatable(luaEventTypeName)
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), eventMethods))

	ud := L.NewUserData()
	ud.Value = event
	L.SetMetatable(ud, L.GetTypeMetatable(luaEventTypeName))

	L.SetGlobal("event", ud)
	return event
}

func (event *Event) Load(L *lua.LState) {
	mt := L.NewTypeMetatable(luaEventTypeName)
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), eventMethods))

	ud := L.NewUserData()
	ud.Value = event
	L.SetMetatable(ud, L.GetTypeMetatable(luaEventTypeName))

	L.SetGlobal("event", ud)
}

func (event *Event) Get(arg string) (ret interface{}) {
	var pointer map[string]interface{}
	var name string
	pointer = event.Values
	names := splitName(arg)

	for i := 0; i < len(names)-1; i++ {
		name = names[i]
		if pointer[name] == nil {
			return nil
		}
		pointer = pointer[name].(map[string]interface{})
	}

	name = names[len(names)-1]
	return pointer[name]
}

func eventGet(L *lua.LState) int {
	var pointer map[string]interface{}
	var name string

	event := checkEvent(L)
	pointer = event.Values

	names := splitName(L.CheckString(2))

	if len(names) == 0 {
		return 0
	}

	for i := 0; i < len(names)-1; i++ {
		name = names[i]
		if pointer[name] == nil {
			return 0
		}
		pointer = pointer[name].(map[string]interface{})
	}
	name = names[len(names)-1]
	switch pointer[name].(type) {
	case string:
		L.Push(lua.LString(pointer[name].(string)))
		return 1
	case float64:
		L.Push(lua.LNumber(pointer[name].(float64)))
		return 1
	case bool:
		L.Push(lua.LBool(pointer[name].(bool)))
		return 1
	default:
		return 0
	}
	return 0
}

func (event *Event) Set(arg string, value interface{}) {
	var pointer map[string]interface{}
	var name string
	if event.Values == nil {
		event.Values = make(map[string]interface{})
	}
	pointer = event.Values
	names := splitName(arg)
	for i := 0; i < len(names)-1; i++ {
		name = names[i]
		if pointer[name] == nil {
			pointer[name] = make(map[string]interface{})
		}
		pointer = pointer[name].(map[string]interface{})
	}
	name = names[len(names)-1]
	pointer[name] = value
}

func eventSet(L *lua.LState) int {
	event := checkEvent(L)

	if event.Values == nil {
		event.Values = make(map[string]interface{})
	}

	names := splitName(L.CheckString(2))

	if len(names) == 0 {
		panic("error")
		return 0
	}

	value := L.CheckAny(3)

	var pointer map[string]interface{}
	pointer = event.Values
	var name string

	for i := 0; i < len(names)-1; i++ {
		name = names[i]
		if pointer[name] == nil {
			pointer[name] = make(map[string]interface{})
		}
		pointer = pointer[name].(map[string]interface{})
	}
	name = names[len(names)-1]
	pointer[name] = mapValue(value)
	return 0
}

func checkEvent(L *lua.LState) (event *Event) {
	ud := L.CheckUserData(1)
	if event, ok := ud.Value.(*Event); ok {
		return event
	}
	return nil
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

func mapValue(value lua.LValue) interface{} {
	switch value.(type) {
	case lua.LString:
		return string(value.(lua.LString))
	case lua.LNumber:
		return float64(value.(lua.LNumber))
	case lua.LBool:
		return bool(value.(lua.LBool))
	case *lua.LTable:
		tb := value.(*lua.LTable)
		len := tb.Len()
		if len == 0 {
			data := make(map[string]interface{})
			tb.ForEach(func(key, val lua.LValue) {
				k := string(key.(lua.LString))
				data[k] = mapValue(val)
			})
			return data
		} else {
			data := make([]interface{}, len)
			tb.ForEach(func(key, val lua.LValue) {
				k := int(key.(lua.LNumber))
				data[k-1] = mapValue(val)
			})
			return data
		}
	default:
		return nil
	}
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

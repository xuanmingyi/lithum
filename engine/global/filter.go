package global

import (
	"engine/libs/gluahttpscrape"
	"engine/libs/gluasql"
	"engine/models"
	"fmt"

	libs "github.com/vadv/gopher-lua-libs"
	lua "github.com/yuin/gopher-lua"
)

type Filter struct {
	Code string
}

func LuaStateFactory() (L *lua.LState) {
	L = lua.NewState()

	// 调用OpenPackage, OpenBase, OpenTable
	for _, pair := range []struct {
		n string
		f lua.LGFunction
	}{
		{lua.LoadLibName, lua.OpenPackage},
		{lua.BaseLibName, lua.OpenBase},
		{lua.TabLibName, lua.OpenTable},
	} {
		if err := L.CallByParam(lua.P{
			Fn:      L.NewFunction(pair.f),
			NRet:    0,
			Protect: true,
		}, lua.LString(pair.n)); err != nil {
			panic(err)
		}
	}

	// 加载 github.com/vadv/gopher-lua-libs
	libs.Preload(L)

	L.SetGlobal("lib_path", lua.LString(Global.LibPath))
	L.SetGlobal("database_driver", lua.LString(Global.Database.Driver))
	L.SetGlobal("database_dsn", lua.LString(Global.Database.DSN))
	L.SetGlobal("debug", lua.LBool(Global.Debug))
	L.SetGlobal("output_path", lua.LString(Global.OutputPath))
	L.SetGlobal("http_proxy", lua.LString(Global.HttpProxy))

	// 其他加载
	gluasql.Preload(L)
	L.PreloadModule("scrape", gluahttpscrape.NewHttpScrapeModule().Loader)

	return L
}

func (f *Filter) Start(input chan models.Event, output chan models.Event) {
	var event models.Event

	for {
		event = <-input
		if len(event.Tags) == 0 {
			L := LuaStateFactory()

			event.Load(L)
			event.Set("body", event.Body)

			err := L.DoString(f.Code)
			if err != nil {
				panic(err)
			}
			output <- event
		} else {
			fmt.Println(event)
		}
	}
}

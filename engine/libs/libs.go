package libs

import (
	"engine/libs/json"
	"engine/libs/gluahttpscrape"
	lua "github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	json.Preload(L)

	L.PreloadModule("scrape", gluahttpscrape.NewHttpScrapeModule().Loader)
}
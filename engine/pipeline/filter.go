package pipeline

import (
	"engine/libs/gluahttpscrape"
	"engine/models"
	"fmt"

	"github.com/tengattack/gluasql"
	libs "github.com/vadv/gopher-lua-libs"
	lua "github.com/yuin/gopher-lua"
)

type Filter struct {
	Code string
}

func (f *Filter) Start(input chan models.Event, output chan models.Event) {
	var event models.Event

	for {
		event = <-input
		if len(event.Tags) == 0 {
			L := lua.NewState()
			libs.Preload(L)

			gluasql.Preload(L)
			L.PreloadModule("scrape", gluahttpscrape.NewHttpScrapeModule().Loader)

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

package outputsqlite

import (
	"context"
	"engine/models"
	"fmt"
)

type OutputSQLite struct {
	Ctx   context.Context
	DSN   string
	Field string
}

const ModuleName = "sqlite"

func InitHandler(ctx context.Context, values map[string]interface{}) (o interface{ Start(chan models.Event) }, err error) {
	output := new(OutputSQLite)
	output.Ctx = ctx
	output.DSN = values["dsn"].(string)
	output.Field = values["field"].(string)
	return output, nil
}

func (t *OutputSQLite) Start(output chan models.Event) {
	for {
		select {
		case <-t.Ctx.Done():
			return
		case event := <-output:
			fmt.Println(event)
		}
	}
}

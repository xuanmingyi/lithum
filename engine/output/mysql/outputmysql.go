package outputmysql

import (
	"context"
	"engine/models"
	"fmt"
)

type Field map[string]string

type OutputMySQL struct {
	Ctx    context.Context
	URL    string
	Table  string
	Fields []interface{}
}

func InitHandler(ctx context.Context, values map[string]interface{}) (output *OutputMySQL, err error) {
	output = new(OutputMySQL)
	output.Ctx = ctx
	output.Table = values["table"].(string)
	output.URL = values["url"].(string)
	output.Fields = values["fields"].([]interface{})
	return output, nil
}

func (t *OutputMySQL) Start(output chan models.Event) {
	for {
		select {
		case <-t.Ctx.Done():
			return
		case event := <-output:
			fmt.Println(event)
		}
	}
}

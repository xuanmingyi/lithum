package outputmysql

import (
	"context"
	"engine/log"
	"engine/models"
	"engine/pipeline"
	"fmt"
	"github.com/sirupsen/logrus"
)

const ModuleName = "mysql"

type Field map[string]string

type OutputMySQL struct {
	Ctx   context.Context
	DSN   string
	Field string
}

func InitHandler(ctx context.Context, values map[string]interface{}) (o pipeline.Output, err error) {
	output := new(OutputMySQL)
	output.Ctx = ctx
	output.DSN = values["dsn"].(string)
	output.Field = values["field"].(string)
	return output, nil
}

func (t *OutputMySQL) Start(output chan models.Event) {
	for {
		select {
		case <-t.Ctx.Done():
			return
		case event := <-output:
			if len(event.Tags) != 0 {
				// 错误
				log.Log.WithFields(logrus.Fields{
					"tags":   event.Tags,
					"module": ModuleName,
				}).Error("数据库错误")
			} else {
				// 运行
				if event.Values[t.Field] != nil {
					sqls := event.Values[t.Field].([]string)
					for _, sql := range sqls {
						t.Ping()
						fmt.Println(sql)
					}
				} else {
					log.Log.WithFields(logrus.Fields{
						"module": ModuleName,
						"field":  t.Field,
					}).Error("配置错误")
				}
			}
		}
	}
}

func (t *OutputMySQL) Ping() {

}

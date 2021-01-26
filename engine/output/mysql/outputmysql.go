package outputmysql

import (
	"context"
	"database/sql"
	"engine/log"
	"engine/models"
	"engine/pipeline"
	"fmt"
	"github.com/sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

const ModuleName = "mysql"

type Field map[string]string

type OutputMySQL struct {
	Ctx   context.Context
	DB *sql.DB
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
						db := t.Ping()
						result, err := db.Exec(sql)
						fmt.Println(result)
						fmt.Println(err)
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

func (t *OutputMySQL) Ping() (DB *sql.DB){
	var err error
	if t.DB != nil {
		if err = t.DB.Ping(); err == nil {
			return t.DB
		}
	}
	t.DB, err = sql.Open("mysql", t.DSN)
	if err != nil {
		panic(err)
	}
	return t.DB
}

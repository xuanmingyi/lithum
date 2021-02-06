package inputmysql

import (
	"context"
	"engine/models"
	"engine/pipeline"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const ModuleName = "mysql"

type InputMySQL struct {
	Ctx      context.Context
	DB       *sql.DB
	Interval int    `yaml:"interval"`
	DSN      string `yaml:"dsn"`
	Code     string `yaml:"code"`
}

func InitHandler(ctx context.Context, values map[string]interface{}) (pipeline.Input, error) {
	input := new(InputMySQL)
	input.Ctx = ctx
	input.Interval = values["interval"].(int)
	input.DSN = values["dsn"].(string)
	input.Code = values["code"].(string)
	return input, nil
}

// 打开数据库
func (model *InputMySQL) Open() (err error) {
	if model.DB, err = sql.Open("mysql", model.DSN); err != nil {
		return err
	}
	if err = model.DB.Ping(); err != nil {
		return err
	}
	return nil
}

// 检查
func (model *InputMySQL) Check() (err error) {
	if model.Interval > 0 {
		return
	}

	if model.DB == nil {
		if err = model.Open(); err != nil {
			return err
		}
	}
	return nil
}

func (model *InputMySQL) Start(msg chan models.Event) {
	startChan := make(chan bool, 1)
	ticker := time.NewTicker(time.Duration(model.Interval) * time.Second)
	defer ticker.Stop()

	startChan <- true

	for {
		select {
		case <-startChan:
			msg <- model.Request()
		case <-ticker.C:
			msg <- model.Request()
		}
	}
}

func (model *InputMySQL) Request() (event models.Event) {
	return event
}

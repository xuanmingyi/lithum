package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type config struct {
	DSN    string  `yaml:"dsn"`
	Models []Model `yaml:"models"`
	DB     *sql.DB `yaml:"-"`
}

var Config config

func init() {
	content, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, &Config)
	if err != nil {
		panic(err)
	}

	err = Config.InitDB()
	if err != nil {
		panic(err)
	}
}

func (c *config) InitDB() (err error) {
	c.DB, err = sql.Open("mysql", c.DSN)
	if err != nil {
		panic(err)
	}
	return nil
}

func (c *config) GetModel(name string) (model Model) {
	for _, model = range c.Models {
		if model.Name == name {
			return model
		}
	}
	return
}

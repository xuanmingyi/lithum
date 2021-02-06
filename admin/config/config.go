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

	for i1, model := range Config.Models {

		for i2, action := range model.Table.TableActions {
			for i3, field := range action.Dialog.Fields {
				Config.Models[i1].Table.TableActions[i2].Dialog.Fields[i3].Attr = model.FindAttrByName(field.Name)
			}
		}

		for i2, action := range model.Table.RowActions {
			for i3, field := range action.Dialog.Fields {
				Config.Models[i1].Table.RowActions[i2].Dialog.Fields[i3].Attr = model.FindAttrByName(field.Name)
			}
		}
	}
}

func (model *Model) FindAttrByName(name string) (attr *Attribute) {
	for _, model := range model.Attributes {
		if model.Name == name {
			return &model
		}
	}
	return nil
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

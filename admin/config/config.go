package config

import (
	"database/sql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type object struct {
	Title string `yaml:"title"`
	Name  string `yaml:"name"`
}

type config struct {
	DSN     string   `yaml:"dsn"`
	Objects []object `yaml:"objects"`

	DB     *sql.DB
	Models []*Model
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

	err = InitDB()
	if err != nil {
		panic(err)
	}

	for _, object := range Config.Objects {
		m, e := LoadModel(object.Name)
		if e != nil {
			panic(e)
		}
		Config.Models = append(Config.Models, m)
	}
}

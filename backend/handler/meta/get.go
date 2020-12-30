package meta

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	. "lithum/handler"
	"lithum/pkg/errno"
	"os"
)

type attribute struct {
	Name    string `yaml:"name" json:"name"`
	Display string `yaml:"display" json:"display"`
}

type model struct {
	Name      string      `yaml:"name" json:"name"`
	Attribute []attribute `yaml:"attributes" json:"attributes"`
}

type column struct {
	Name string `yaml:"name" json:"name"`
}

type field struct {
	Name string `yaml:"name" json:"name"`
}

type dialog struct {
	Title  string  `yaml:"title" json:"title"`
	Fields []field `yaml: "fields" json: "fields"`
}

type action struct {
	Name    string `yaml:"name" json:"name"`
	Display string `yaml:"display" json:"display"`
	Type    string `yaml:"type" json:"type"`
	Dialog  dialog `yaml:"dialog" json:"dialog"`
}

type table struct {
	Columns    []column `yaml:"columns" json:"columns"`
	Actions    []action `yaml:"actions" json:"actions"`
	RowActions []action `yaml:"row_actions" json:"row_actions"`
}

type metaData struct {
	Model model `yaml:"model" json:"model"`
	Table table `json:"table" json:"table"`
}

func Get(c *gin.Context) {
	var meta metaData

	name := c.Param("name")
	path := fmt.Sprintf("conf/meta/%v.yaml", name)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		SendResponse(c, errno.ErrMetaNotFound, nil)
		return
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		SendResponse(c, errno.ErrMetaRead, nil)
		return
	}

	err = yaml.Unmarshal(content, &meta)
	if err != nil {
		SendResponse(c, errno.ErrMetaYaml, nil)
		return
	}

	SendResponse(c, nil, meta)
}

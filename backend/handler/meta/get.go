package meta

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	. "lithum/handler"
	. "lithum/model"
	"lithum/pkg/errno"
	"os"
)

func Get(c *gin.Context) {
	var meta MetaData

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

package form

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"lithum/handler"
)

func FormStruct(c *gin.Context) {
	data, err := ioutil.ReadFile("a.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	handler.SendResponse(c, nil, string(data))
}

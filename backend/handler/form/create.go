package form

import (
	"github.com/gin-gonic/gin"
	"lithum/handler"
)

func Create(c *gin.Context) {
	//	data, err := ioutil.ReadFile("a.json")
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(string(data))
	//	handler.SendResponse(c, nil, string(data))

	handler.SendResponse(c, nil, nil)
}

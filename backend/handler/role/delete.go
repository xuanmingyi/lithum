package role

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "lithum/handler"
)

func Delete(c *gin.Context) {
	fmt.Println("delete1!!!")
	SendResponse(c, nil, nil)
}

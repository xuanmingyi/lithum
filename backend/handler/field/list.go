package field

import (
	"github.com/gin-gonic/gin"
	. "lithum/handler"
)

func List(c *gin.Context) {
	SendResponse(c, nil, nil)
}

package field

import (
	"github.com/gin-gonic/gin"
	. "lithum/handler"
)

func Delete(c *gin.Context) {
	SendResponse(c, nil, nil)
}

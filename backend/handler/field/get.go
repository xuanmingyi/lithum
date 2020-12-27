package field

import (
	"github.com/gin-gonic/gin"
	. "lithum/handler"
)

func Get(c *gin.Context) {
	SendResponse(c, nil, nil)
}

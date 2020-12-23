package form

import (
	"github.com/gin-gonic/gin"
	"lithum/handler"
)

func Get(c *gin.Context) {
	handler.SendResponse(c, nil, nil)
}
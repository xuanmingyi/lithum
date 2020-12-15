package user

import (
	"github.com/gin-gonic/gin"
	"lithum/handler"
)

func Logout(c *gin.Context) {
	handler.SendResponse(c, nil, nil)
}

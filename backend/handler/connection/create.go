package connection

import (
	"github.com/gin-gonic/gin"

	"lithum/model"
	"lithum/pkg/errno"

	. "lithum/handler"
	. "lithum/model/request/connection"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	con := model.Connection{
		Name:    r.Name,
		Display: r.Display,
		Config:  r.Config,
	}

	if err := con.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}

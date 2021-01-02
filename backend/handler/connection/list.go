package connection

import (
	"github.com/gin-gonic/gin"
	. "lithum/handler"
	"lithum/model"
	. "lithum/model/request/connection"
	. "lithum/model/response/connection"
	"lithum/pkg/errno"
)

func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	offset := (r.Page - 1) * r.Limit
	items, count, err := model.ListConnection(r.Search, offset, r.Limit)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, ListResponse{
		Count: count,
		Items: items,
	})
}

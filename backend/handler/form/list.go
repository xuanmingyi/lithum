package form

import (
	"github.com/gin-gonic/gin"
	. "lithum/handler"
	. "lithum/model/request/form"
	. "lithum/model/response/form"
	"lithum/pkg/errno"
	"lithum/service"
)

func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	offset := (r.Page - 1) * r.Limit
	items, count, err := service.ListForm(r.Search, offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		Count: count,
		Items: items,
	})
}

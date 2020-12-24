package field

import (
	"github.com/gin-gonic/gin"
	"lithum/handler"
	"lithum/model"
	. "lithum/model/request/form"
	. "lithum/model/response/form"
	"lithum/pkg/errno"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	f := model.SysForm{
		Size:          r.Size,
		LabelPosition: r.LabelPosition,
		LabelWidth:    r.LabelWidth,
		Span:          r.Span,
		FormBtns:      r.FormBtns,
		Disabled:      r.Disabled,
	}

	if err := f.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
	}
	handler.SendResponse(c, nil, nil)
}

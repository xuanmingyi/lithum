package form

import (
	"github.com/gin-gonic/gin"
	. "lithum/handler"
	"lithum/model"
	. "lithum/model/request/form"
	. "lithum/model/response/form"
	"lithum/pkg/errno"
)

func Create(c *gin.Context) {
	/*
		data, err := ioutil.ReadFile("a.json")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))

		handler.SendResponse(c, nil, string(data))
		//fmt.Println("hello world")
		//handler.SendResponse(c, nil, nil)

	*/
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	f := model.SysForm{
		Title:         r.Title,
		LabelWidth:    r.LabelWidth,
		LabelPosition: r.LabelPosition,
		Size:          r.Size,
		Span:          r.Span,
		FormBtns:      r.FormBtns,
		Disabled:      r.Disabled,
	}

	if err := f.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		ID: f.ID,
	}
	SendResponse(c, nil, rsp)
}

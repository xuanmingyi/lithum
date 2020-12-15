package user

import (
	. "lithum/handler"
	. "lithum/model/response/user"

	"lithum/model"
	"lithum/pkg/errno"
	"lithum/pkg/token"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {

	ctx, err := token.ParseRequest(c)
	if err != nil {
		SendResponse(c, errno.ErrTokenInvalid, nil)
		return
	}

	d, err := model.GetUser(ctx.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, MeResponse{
		Username:     d.Username,
		Nickname:     d.Nickname,
		Introduction: d.Introduction,
		Avatar:       d.Avatar,
	})
}

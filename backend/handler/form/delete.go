package form

import (
	"github.com/gin-gonic/gin"
	. "lithum/handler"
	"lithum/model"
	"strconv"
)

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteForm(uint64(id)); err != nil {
		SendResponse(c, err, nil)
	}
	SendResponse(c, nil, nil)
}

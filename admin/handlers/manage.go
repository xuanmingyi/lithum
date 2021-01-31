package handlers

import (
	"admin/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


// url:  /manage/:name/index
// desc: 首页
func ManageHandler(c *gin.Context) {
	name := c.Param("name")
	var current *config.Model

	for _, m := range config.Config.Models {
		if m.Name == name {
			current = m
		}
	}

	fmt.Println(name)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Gin 模板",
		"models": config.Config.Models,
		"current": current,
	})
}


// url:  /manage/:name/create
// desc: 创建
func CreateHandler(c *gin.Context) {
	name := c.Param("name")

	fmt.Println(name)

	c.HTML(http.StatusOK, "form.html", gin.H{
		"title": "ssss",
	})
}


// url:  /manage/:name/data
// desc: 获取数据
func DataHandler(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
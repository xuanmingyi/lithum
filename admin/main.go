package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"admin/config"
)

func main() {

	fmt.Println(config.Config)

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Gin 模板",
		})
	})
	r.Run()
}

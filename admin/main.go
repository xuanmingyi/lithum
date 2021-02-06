package main

import (
	"admin/config"
	"admin/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/manage/:name/index", handlers.ManageHandler)
	router.GET("/manage/:name/table", handlers.TableHandler)
	router.GET("/manage/:name/form", handlers.FormHandler)
	router.GET("/manage/:name/data", handlers.DataHandler)
	router.GET("/", func(c *gin.Context) {
		if len(config.Config.Models) >= 1 {
			fmt.Println(config.Config.Models[0].Name)
			c.Request.URL.Path = "/manage/" + config.Config.Models[0].Name + "/index"
		}
		router.HandleContext(c)
	})

	router.Run("0.0.0.0:8081")
}

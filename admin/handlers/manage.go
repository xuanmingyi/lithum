package handlers

import (
	"admin/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// url:  /manage/:name/index
// desc: 首页
func ManageHandler(c *gin.Context) {
	name := c.Param("name")
	current_model := config.Config.GetModel(name)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":         "Gin 模板",
		"models":        config.Config.Models,
		"current_model": current_model,
	})
}

// url:  /manage/:name/create
// desc: 创建
func CreateHandler(c *gin.Context) {
	name := c.Param("name")
	current_model := config.Config.GetModel(name)

	c.HTML(http.StatusOK, "form.html", gin.H{
		"title":         "ssss",
		"current_model": current_model,
	})
}

// url:  /manage/:name/data
// desc: 获取数据
func DataHandler(c *gin.Context) {
	name := c.Param("name")
	page, err := strconv.Atoi(c.Query("page"))
	limit, err := strconv.Atoi(c.Query("limit"))
	ret := make(map[string]interface{})
	var count int64


	count_row, err := config.Config.DB.Query(fmt.Sprintf("SELECT count(*) FROM %s", name))
	if err != nil {
		panic(err)
	}

	for count_row.Next(){
		err = count_row.Scan(&count)
		if err != nil {
			panic(err)
		}
		fmt.Println(count)
	}

	rows, err := config.Config.DB.Query(fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC LIMIT %d OFFSET %d", name, limit, (page - 1)*limit))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	columns, _ := rows.Columns()
	column_length := len(columns)

	ret["code"] = 0
	ret["msg"] = ""
	ret["count"] = count
	ret["data"] = make([]map[string]interface{}, 0)

	line_cache := make([]interface{}, column_length)
	for index, _ := range line_cache {
		var i interface{}
		line_cache[index] = &i
	}

	for rows.Next() {
		rows.Scan(line_cache...)
		item := make(map[string]interface{})
		for index, value := range line_cache {
			item[columns[index]] = *value.(*interface{})
		}
		ret["data"] = append(ret["data"].([]map[string]interface{}), item)
	}

	c.JSON(http.StatusOK, ret)
}

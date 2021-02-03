package handlers

import (
	"admin/config"
	"database/sql"
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

func TableHandler(c *gin.Context) {
	name := c.Param("name")
	current_model := config.Config.GetModel(name)
	c.HTML(http.StatusOK, "table.html", gin.H{
		"title": "sss",
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

	values := make([]sql.RawBytes, column_length)
	line_cache := make([]interface{}, column_length)
	for index, _ := range line_cache {
		line_cache[index] = &values[index]
	}

	// https://docs.hacknode.org/gopl-zh/ch7/ch7-05.html
	// https://docs.hacknode.org/gopl-zh/ch12/ch12-05.html
	// https://studygolang.com/articles/3244
	for rows.Next() {
		rows.Scan(line_cache...)
		item := make(map[string]interface{})
		for index, value := range values {
			item[columns[index]] = string(value)
		}
		ret["data"] = append(ret["data"].([]map[string]interface{}), item)
	}

	c.JSON(http.StatusOK, ret)
}
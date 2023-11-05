package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ArticleModel struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// _getList 文章列表页面
func _getList(c *gin.Context) {

	articleList := []ArticleModel{
		{"上课", "震惊"},
		{"女神", "震惊"},
		{"老师", "震惊"},
	}

	c.JSON(200, Response{0, articleList, "请求成功"})
}

// _getDetail 文章详情页面
func _getDetail(c *gin.Context) {
	// 获取param 的id
	fmt.Println(c.Param("id"))
	result := ArticleModel{"上课", "震惊"}
	c.JSON(200, Response{0, result, "查询成功"})
}

// _create 文章创建页面
func _create(c *gin.Context) {
	//接收前端的JSON
	var articles ArticleModel

	err := _bindJson(c, &articles)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, Response{0, articles, "添加成功"})
}

// _update 编辑文章
func _update(c *gin.Context) {
	//接收id
	fmt.Println(c.Param("id"))

	//接收前端的JSON
	var articles ArticleModel

	err := _bindJson(c, &articles)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, Response{0, articles, "修改成功"})
}

func _delete(c *gin.Context) {

	//接收id
	id := c.Param("id")
	fmt.Println(id)
	c.JSON(200, Response{0, id, "删除成功"})
}

func main() {
	router := gin.Default()

	// Create a router for the articles endpoint
	router.GET("/articles", _getList)
	// Create a GET request for the articles endpoint
	router.GET("/articles/:id", _getDetail)
	// Create a POST request for the articles endpoint
	router.POST("/articles", _create)
	// Create a PUT request for the articles endpoint
	router.PUT("/articles/:id", _update)
	// Create a DELETE request for the articles endpoint
	router.DELETE("/articles/:id", _delete)
	router.Run(":80")
}

func _bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

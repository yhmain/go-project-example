package main

import (
	"fmt"
	"os"

	cotroller "github.com/yhmain/go-project-example/controller"
	"github.com/yhmain/go-project-example/repository"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	if err := Init("./data/"); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()

	//测试链接：localhost:8080/community/page/get/2
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := cotroller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	// 新增发帖功能
	// 需要传入2个参数：该帖对应的主题ID，和该帖的内容
	r.POST("/community/post/submit", func(ctx *gin.Context) {
		topicId := ctx.PostForm("topicId")
		postContent := ctx.PostForm("content")
		// data := cotroller.PublishPost(topicId, content)
		// ctx.JSON(200, data)
		fmt.Printf("ID: %v, Content: %v", topicId, postContent)
	})
	err := r.Run()
	if err != nil {
		return
	}
}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		return err
	}
	return nil
}

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/someGet", func(c *gin.Context) {

	})
	//r.POST("/upload", func(c *gin.Context) {
	//	// 单文件
	//	file, _ := c.FormFile("file")
	//	log.Println(file.Filename)
	//
	//	// 上传文件到指定的 dst 。
	//	// c.SaveUploadedFile(file, dst)
	//
	//	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	//})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 默认情况下，它使用：8080，除非定义了 PORT 环境变量。
	r.Run() // 在 127.0.0.1:8080 上监听并服务
}

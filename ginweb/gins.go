package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @program:   go-basic-exercises
// @file:      gins.go
// @author:    bowen
// @time:      2023-09-16 17:38
// @description:

func main() {
	server := gin.Default()
	server.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello go")
	})
	go func() {
		server1 := gin.Default()
		server1.GET("/hell2", func(c *gin.Context) {
			c.String(http.StatusOK, "hello go")
		})
		server1.Run(":8081")
	}()

	server.POST("/post", func(c *gin.Context) {
		c.String(200, "这是post请求")
	})

	//参数路由
	server.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, "这是参数路由"+name)
	})

	//匹配符路由
	server.GET("/views/*.html", func(c *gin.Context) {
		path := c.Param(".html")
		c.String(200, "这是通配符路由 %s", path)
	})

	//查询参数
	server.GET("/order", func(c *gin.Context) {
		oid := c.Query("id")
		c.JSON(http.StatusOK, oid)
	})

	server.Run(":8080")
}

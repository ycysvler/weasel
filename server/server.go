package server

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	// 设置 gin 的模式（调试模式：DebugMode, 发行模式：ReleaseMode）
	gin.SetMode(gin.DebugMode)
	// 创建一个不包含中间件的路由器
	r := gin.Default()

	r.Static("/static", "./static")
	// static icon
	r.StaticFile("/favicon.ico", "./static/images/favicon.ico")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/graphql", graphiql)

	r.POST("/api", api)

	// 在8080 端口，启动http服务
	r.Run(":8090")

}

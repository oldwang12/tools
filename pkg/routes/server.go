package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/oldwang12/tools/pkg/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 简单的路由组: v1
	v1 := r.Group("/api/v1")
	v1.Use(PrintIP)
	{
		base64 := v1.Group("/tools")
		base64.POST("/base64", controllers.Base64Msg)
		base64.GET("/ip", controllers.QueryCidr)
	}

	return r
}

func PrintIP(c *gin.Context) {
	if c.ClientIP() == "::1" {
		log.Print("127.0.0.1", ",url: ", c.Request.URL)
		return
	}
	log.Print(c.ClientIP(), ",url: ", c.Request.URL)
}

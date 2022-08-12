package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oldwang12/tools/pkg/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 简单的路由组: v1
	v1 := r.Group("/api/v1")

	{
		base64 := v1.Group("/tools/base64")
		base64.POST("/", controllers.Base64Msg)
	}

	return r
}

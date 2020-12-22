package routers

import (
	"gin-web/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api")
	{
		//获取user列表
		apiv1.GET("/users/:id", api.GetUser)
	}

	return r
}

package router

import (
	"github.com/gin-gonic/gin"
	"web/controller"
)

func TestRoute(g *gin.RouterGroup) {
	//登录
	group := g.Group("/test")
	{
		group.GET("/fang", controller.Fang)
	}
}

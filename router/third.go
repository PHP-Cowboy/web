package router

import (
	"github.com/gin-gonic/gin"
	"web/controller"
)

func ThirdRoute(g *gin.RouterGroup) {
	//登录
	group := g.Group("/third")
	{
		group.POST("/wx", controller.Wx)
	}

}

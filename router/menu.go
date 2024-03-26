package router

import (
	"github.com/gin-gonic/gin"
	"web/controller"
	"web/middlewares"
)

func MenuRoute(g *gin.RouterGroup) {

	menuGroup := g.Group("/menu", middlewares.JWTAuth())
	{
		//菜单列表
		menuGroup.GET("/list", controller.GetMenuList)
		//新增菜单
		menuGroup.POST("/create", controller.CreateMenu)
		//修改菜单
		menuGroup.POST("/change", controller.ChangeMenu)
		//批量删除菜单
		menuGroup.POST("/batch_delete", controller.BatchDeleteMenu)

	}
}

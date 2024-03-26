package router

import (
	"github.com/gin-gonic/gin"
	"web/controller"
	"web/middlewares"
)

func RoleRoute(g *gin.RouterGroup) {
	roleGroup := g.Group("/role", middlewares.JWTAuth())
	{
		//新增角色
		roleGroup.POST("/create", controller.CreateRole)
		//角色列表
		roleGroup.GET("/list", controller.GetRoleList)
		//批量删除角色
		roleGroup.POST("/batch_delete", controller.BatchDeleteRole)
	}
}

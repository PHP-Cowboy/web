package router

import (
	"github.com/gin-gonic/gin"
	"web/controller"
	"web/middlewares"
)

func UserRoute(g *gin.RouterGroup) {
	//登录
	group := g.Group("/user")
	{
		group.POST("/login", controller.Login)
		//
		group.GET("/getUser", controller.GetUser)
	}

	loginGroup := g.Group("/user", middlewares.JWTAuth())
	{
		//校验登录是否过期
		loginGroup.GET("/check", controller.Check)
	}

	userGroup := g.Group("/user", middlewares.JWTAuth(), middlewares.IsSuperAdminAuth())
	{
		//用户列表
		userGroup.GET("/list", controller.GetUserList)
		//新增用户
		userGroup.POST("/create", controller.CreateUser)
		//修改密码
		userGroup.POST("/change", controller.ChangeUser)
		//批量删除用户
		userGroup.POST("/batch_delete", controller.BatchDeleteUser)

	}
}

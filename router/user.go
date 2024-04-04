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
		group.POST("/loginByPwd", controller.LoginByPwd)
		group.POST("/loginByCode", controller.LoginByCode)
		group.POST("/registration", controller.Registration)
	}

	loginGroup := g.Group("/user", middlewares.JWTAuth())
	{
		//校验登录是否过期
		loginGroup.GET("/check", controller.Check)
		//修改密码
		loginGroup.POST("/change", controller.ChangeUser)
		//建议
		loginGroup.POST("/suggestion", controller.Suggestion)
	}

}

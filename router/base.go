package router

import (
	"github.com/gin-gonic/gin"
	"web/controller"
)

func BaseRoute(g *gin.RouterGroup) {
	baseGroup := g.Group("/base")
	{
		//验证码
		baseGroup.GET("/captcha", controller.GetCaptcha)
	}

	test := g.Group("/test")

	{
		test.POST("/t", controller.Test)
	}
}

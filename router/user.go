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
		group.POST("/wxLogin", controller.WxLogin)
	}

	loginGroup := g.Group("/user", middlewares.JWTAuth())
	{
		//校验登录是否过期
		loginGroup.GET("/check", controller.Check)
		//修改密码
		loginGroup.POST("/change", controller.ChangeUser)
		//建议
		loginGroup.POST("/suggestion", controller.Suggestion)
		//支付礼包列表
		loginGroup.POST("/payGiftList", controller.PayGiftList)
		//根据礼包创建订单
		loginGroup.POST("/order", controller.Order)
	}

}

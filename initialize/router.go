package initialize

import (
	"github.com/gin-gonic/gin"
	"web/middlewares"
	"web/router"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	//跨域
	r.Use(middlewares.Cors())

	group := r.Group("/v1")
	//base
	router.BaseRoute(group)
	//用户
	router.UserRoute(group)
	//角色
	router.RoleRoute(group)
	//菜单
	router.MenuRoute(group)
	//字典
	router.DictRoute(group)
	//测试
	router.TestRoute(group)

	return r
}

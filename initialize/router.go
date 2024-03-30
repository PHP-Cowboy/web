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
	//医案
	router.MedicalCasesRoute(group)
	//大数据
	router.HomeRoute(group)
	//用户
	router.UserRoute(group)

	return r
}

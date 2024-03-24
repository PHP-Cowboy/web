package router

import (
	"github.com/gin-gonic/gin"
	"web/controller"
)

func BigDataRoute(g *gin.RouterGroup) {
	//医疗案例
	bg := g.Group("/bigData")
	{
		bg.POST("/list", controller.MedicalCasesList)     //列表
		bg.POST("/detail", controller.MedicalCasesDetail) //详情
	}
}

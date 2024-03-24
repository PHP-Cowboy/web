package router

import (
	"github.com/gin-gonic/gin"
	"web/controller"
)

func MedicalCasesRoute(g *gin.RouterGroup) {
	//医疗案例
	mg := g.Group("/medicalCases")
	{
		mg.POST("/list", controller.MedicalCasesList)     //列表
		mg.POST("/detail", controller.MedicalCasesDetail) //详情
	}
}

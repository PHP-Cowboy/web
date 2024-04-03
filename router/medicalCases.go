package router

import (
	"github.com/gin-gonic/gin"
	"web/controller"
	"web/middlewares"
)

func MedicalCasesRoute(g *gin.RouterGroup) {
	//医疗案例
	mg := g.Group("/medicalCases", middlewares.JWTAuth())
	{
		mg.POST("/list", controller.MedicalCasesList)         //列表
		mg.POST("/randList", controller.RandMedicalCasesList) //列表
		mg.POST("/detail", controller.MedicalCasesDetail)     //详情
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"web/controller"
)

func MedicalCasesRoute(g *gin.RouterGroup) {
	//医疗案例
	medicalCasesGroup := g.Group("/medicalCases")
	{
		//列表
		medicalCasesGroup.GET("/list", controller.MedicalCasesList)
		medicalCasesGroup.POST("/detail", controller.MedicalCasesDetail)
	}
}

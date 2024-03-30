package router

import (
	"github.com/gin-gonic/gin"
	"web/controller"
	"web/middlewares"
)

func HomeRoute(g *gin.RouterGroup) {
	//医疗案例
	h := g.Group("/home", middlewares.JWTAuth())
	bg := h.Group("/bigData")
	{
		bg.POST("/list", controller.BigDataList)          //列表
		bg.POST("/detail", controller.MedicalCasesDetail) //详情
	}
	cs := h.Group("/classics")
	{
		cs.POST("/categoryList", controller.ClassicsCategoryList) //典籍分类列表
		cs.POST("/list", controller.BookListByCategory)           //某个分类的典籍列表
		cs.POST("/catalogueList", controller.CatalogueList)       //典籍目录列表
		cs.POST("/bookContent", controller.BookContent)           //典籍的内容
		cs.POST("/prevBookContent", controller.PrevBookContent)   //上一章
		cs.POST("/nextBookContent", controller.NextBookContent)   //下一章
	}
	t := h.Group("/tool")
	{
		t.POST("/list", controller.ToolList) //工具列表
	}
}

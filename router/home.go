package router

import (
	"github.com/gin-gonic/gin"
	"web/controller"
	"web/middlewares"
)

func HomeRoute(g *gin.RouterGroup) {
	//医疗案例
	mg := g.Group("/medicalCases", middlewares.JWTAuth())
	{
		mg.POST("/list", controller.MedicalCasesList)         //列表
		mg.POST("/randList", controller.RandMedicalCasesList) //列表
		mg.POST("/detail", controller.MedicalCasesDetail)     //详情
	}
	//首页
	h := g.Group("/home", middlewares.JWTAuth())

	//方剂
	p := h.Group("/prescription")
	{
		p.POST("/randList", controller.PrescriptionRandList)                                           //大数据随机获取列表
		p.POST("/list", controller.PrescriptionList)                                                   //大数据方剂列表
		p.POST("/info", controller.PrescriptionInfo)                                                   //大数据方剂详情
		p.POST("/celebrityList", controller.PrescriptionCelebrityList)                                 //大数据方剂名家列表
		p.POST("/celebrityInfo", controller.PrescriptionCelebrityInfo)                                 //大数据方剂名家详情
		p.POST("/celebrityInfoByPrescriptionId", controller.PrescriptionCelebrityInfoByPrescriptionId) //根据方剂id获取名家详情
		p.POST("/graph", controller.PrescriptionGraph)                                                 //大数据方剂名家列表
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
	m := h.Group("/mind")
	{
		m.POST("/list", controller.MindMapList) //思维导图列表
	}
	//常见疾病辩证辅助
	d := h.Group("/disease")
	{
		d.POST("/category", controller.DiseaseCategory) //疾病分类
		d.POST("/disease", controller.DiseaseList)      //疾病分类数据
	}
	//名医心法
	mm := h.Group("/mindMethod")
	{
		mm.POST("/list", controller.MindMethodList) //列表
		mm.POST("/info", controller.MindMethod)     //详情
	}
	//常用方剂
	cp := h.Group("commonlyPrescription")
	{
		cp.POST("/categoryList", controller.CommonlyPrescriptionCategoryList) //分类列表
		cp.POST("/list", controller.CommonlyPrescriptionList)                 //列表
		cp.POST("/info", controller.CommonlyPrescription)                     //详情
	}
	//方剂大全
	ccp := h.Group("completeCollectionPrescription")
	{
		ccp.POST("/list", controller.CompleteCollectionPrescriptionList) //列表
		ccp.POST("/info", controller.CompleteCollectionPrescription)     //详情
	}
}

package controller

import (
	"github.com/gin-gonic/gin"
	"web/daos"
	"web/forms/req"
	"web/utils/ecode"
	"web/utils/xsq_net"
)

func ToolList(c *gin.Context) {
	form := req.ToolList{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.ToolList(form)
	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 思维导图列表
func MindMapList(c *gin.Context) {
	form := req.MindMapList{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.MindMapList(form)
	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 疾病分类
func DiseaseCategory(c *gin.Context) {

	res, err := daos.DiseaseCategory()
	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 疾病分类数据
func DiseaseList(c *gin.Context) {
	form := req.DiseaseList{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.Disease(form)
	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 名医心法列表
func MindMethodList(c *gin.Context) {
	form := req.MindMethodList{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.MindMethodList(form)
	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 名医心法详情
func MindMethod(c *gin.Context) {
	form := req.Id{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.MindMethod(form)
	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 中医方剂分类列表
func CommonlyPrescriptionCategoryList(c *gin.Context) {
	form := req.CommonlyPrescriptionCategoryList{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.CommonlyPrescriptionCategoryList(form)
	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 常用方剂列表
func CommonlyPrescriptionList(c *gin.Context) {
	form := req.CommonlyPrescriptionList{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.CommonlyPrescriptionList(form)
	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 常用方剂详情
func CommonlyPrescription(c *gin.Context) {
	form := req.Id{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.CommonlyPrescription(form)
	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 方剂大全列表
func CompleteCollectionPrescriptionList(c *gin.Context) {
	form := req.CompleteCollectionPrescriptionList{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.CompleteCollectionPrescriptionList(form)
	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 方剂大全详情
func CompleteCollectionPrescription(c *gin.Context) {
	form := req.Id{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.CompleteCollectionPrescription(form)
	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

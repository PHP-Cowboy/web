package controller

import (
	"github.com/gin-gonic/gin"
	"web/daos"
	"web/forms/req"
	"web/utils/ecode"
	"web/utils/xsq_net"
)

// 大数据随机获取列表
func PrescriptionRandList(c *gin.Context) {
	form := req.PrescriptionRandList{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.PrescriptionRandList(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 大数据方剂列表
func PrescriptionList(c *gin.Context) {
	form := req.PrescriptionList{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.PrescriptionList(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 大数据方剂详情
func PrescriptionInfo(c *gin.Context) {
	form := req.Id{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.PrescriptionInfo(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 大数据方剂名家列表
func PrescriptionCelebrityList(c *gin.Context) {
	form := req.PrescriptionCelebrityList{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.PrescriptionCelebrityList(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 大数据方剂名家详情
func PrescriptionCelebrityInfo(c *gin.Context) {
	form := req.Id{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.PrescriptionCelebrityInfo(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 根据方剂id获取大数据方剂名家详情
func PrescriptionCelebrityInfoByPrescriptionId(c *gin.Context) {
	form := req.Id{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.PrescriptionCelebrityInfoByPrescriptionId(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 大数据方剂名家列表
func PrescriptionGraph(c *gin.Context) {
	form := req.PrescriptionGraph{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.PrescriptionGraph(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

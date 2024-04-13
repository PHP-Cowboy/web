package controller

import (
	"github.com/gin-gonic/gin"
	"web/daos"
	"web/forms/req"
	"web/utils/ecode"
	"web/utils/xsq_net"
)

// 列表
func ClinicalList(c *gin.Context) {
	form := req.ClinicalList{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.ClinicalList(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 大数据饼图
func BigDataPieChart(c *gin.Context) {
	form := req.BigDataPieChart{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.BigDataPieChart(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 大数据饼图
func RandOne(c *gin.Context) {
	form := req.BigDataPieChart{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.BigDataPieChart(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

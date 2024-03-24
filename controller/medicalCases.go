package controller

import (
	"github.com/gin-gonic/gin"
	"web/daos"
	"web/forms/req"
	"web/utils/ecode"
	"web/utils/xsq_net"
)

// 医疗案例列表
func MedicalCasesList(c *gin.Context) {
	var form req.MedicalCasesList

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.MedicalCasesList(form)
	if err != nil {
		return
	}

	xsq_net.SucJson(c, res)
}

// 医疗案例详情
func MedicalCasesDetail(c *gin.Context) {
	var form req.Id

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.MedicalCasesDetail(form)
	if err != nil {
		return
	}

	xsq_net.SucJson(c, res)
}

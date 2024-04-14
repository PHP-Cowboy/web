package controller

import (
	"github.com/gin-gonic/gin"
	"web/daos"
	"web/forms/req"
	"web/utils/ecode"
	"web/utils/xsq_net"
)

// 伤寒论目录列表
func TyphoidCatalogueList(c *gin.Context) {
	form := req.TyphoidCatalogueList{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.TyphoidCatalogueList(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 伤寒论的内容
func TyphoidInfo(c *gin.Context) {
	form := req.Id{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.TyphoidInfo(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 伤寒论上一章
func TyphoidPrev(c *gin.Context) {
	form := req.Id{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.TyphoidPrev(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 伤寒论下一章
func TyphoidNext(c *gin.Context) {
	form := req.Id{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.TyphoidNext(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

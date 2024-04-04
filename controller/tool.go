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
		return
	}

	xsq_net.SucJson(c, res)
}

func MindMapList(c *gin.Context) {
	form := req.MindMapList{}

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	res, err := daos.MindMapList(form)
	if err != nil {
		return
	}

	xsq_net.SucJson(c, res)
}

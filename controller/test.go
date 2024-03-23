package controller

import (
	"github.com/gin-gonic/gin"
	"strings"
	"web/global"
	"web/model"
	"web/utils/ecode"
	"web/utils/xsq_net"
)

type FangReq struct {
	Mingcheng string `json:"mingcheng" form:"mingcheng"`
}

func Fang(c *gin.Context) {
	var req FangReq

	if err := c.ShouldBind(&req); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	var fang model.Fang

	db := global.DB

	result := db.Model(&model.Fang{}).Where(&model.Fang{Mingcheng: strings.Trim(req.Mingcheng, " ")}).First(&fang)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	xsq_net.SucJson(c, fang)
}

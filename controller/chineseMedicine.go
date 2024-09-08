package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"web/daos"
	"web/forms/req"
	"web/utils/ecode"
	"web/utils/xsq_net"
)

// 中药列表
func ChineseMedicineList(c *gin.Context) {
	var form req.MedicineList

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	uid, ok := c.Get("uid")

	if !ok {
		xsq_net.ErrorJSON(c, ecode.UserNotLogin)
		return
	}

	form.UserId = uid.(int)

	res, err := daos.MedicineList(form)
	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, res)
}

// 新增中药
func AddChineseMedicine(c *gin.Context) {
	var form req.SaveMedicine

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	err := daos.SaveMedicine(form)
	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.Success(c)
}

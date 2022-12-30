package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"time"
	"web/forms/req"
	"web/forms/rsp"
	"web/global"
	"web/model"
	"web/utils/ecode"
	"web/utils/timeutil"
	"web/utils/xsq_net"
)

// 创建角色
func CreateRole(c *gin.Context) {
	var form req.CreateRoleForm

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	result := global.DB.Create(&model.Role{Name: form.Name})

	if result.Error != nil || result.RowsAffected == 0 {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	xsq_net.Success(c)
}

// 修改角色
func ChangeRole(c *gin.Context) {
	var form req.ChangeRoleForm

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	db := global.DB

	var role model.Role

	result := db.First(&role, form.Id)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	if result.RowsAffected == 0 {
		xsq_net.ErrorJSON(c, ecode.RoleNotFound)
		return
	}

	deleteTime, err := time.ParseInLocation(timeutil.TimeFormat, timeutil.GetDateTime(), time.Local)
	if err != nil {
		xsq_net.ErrorJSON(c, ecode.DataTransformationError)
		return
	}

	if form.IsDelete {
		result = db.Model(&role).Updates(model.Role{
			Base: model.Base{
				DeleteTime: deleteTime,
			},
		})
	}

	xsq_net.Success(c)
}

// 角色列表
func GetRoleList(c *gin.Context) {
	var form req.GetRoleListForm

	err := c.ShouldBind(&form)

	if err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	db := global.DB

	var (
		roles []model.Role
		res   rsp.GetRoleList
	)

	result := db.Where("delete_time is null").Find(&roles)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	res.Total = result.RowsAffected

	db.Where("delete_time is null").Scopes(model.Paginate(form.Page, form.Size)).Find(&roles)

	for _, role := range roles {
		res.List = append(res.List, &rsp.Role{
			Id:         role.Id,
			CreateTime: role.CreateTime.Format(timeutil.TimeFormat),
			Name:       role.Name,
		})
	}

	xsq_net.SucJson(c, res)
}

// 批量删除角色
func BatchDeleteRole(c *gin.Context) {
	var form req.BatchDeleteRoleForm

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	result := global.DB.Model(model.Role{}).
		Where("id in (?)", form.Ids).
		Updates(map[string]interface{}{
			"delete_time": time.Now().Format(timeutil.TimeFormat),
		})

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	xsq_net.Success(c)
}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strings"
	"time"
	"web/forms/req"
	"web/forms/rsp"
	"web/global"
	"web/model"
	"web/utils/ecode"
	"web/utils/timeutil"
	"web/utils/xsq_net"
)

// 字典类型列表
func DictTypeList(c *gin.Context) {
	var form req.DictTypeListForm

	err := c.ShouldBind(&form)
	if err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	var (
		types []model.DictType
		res   rsp.DictTypeListRsp
	)

	db := global.DB

	result := db.Where(model.DictType{Code: form.Code, Name: form.Name}).Where("delete_time is null").Find(&types)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	res.Total = result.RowsAffected

	db.Where(model.DictType{Code: form.Code, Name: form.Name}).Where("delete_time is null").Scopes(model.Paginate(form.Page, form.Size)).Find(&types)

	list := make([]*rsp.DictType, 0)
	for _, t := range types {
		list = append(list, &rsp.DictType{
			Code:       t.Code,
			Name:       t.Name,
			CreateTime: t.CreateTime.Format(timeutil.TimeFormat),
		})
	}

	res.List = list

	xsq_net.SucJson(c, res)
}

// 新增字典类型
func CreateDictType(c *gin.Context) {
	var form req.CreateDictTypeForm

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	db := global.DB

	dict := model.DictType{
		Code: strings.ToLower(form.Code),
		Name: form.Name,
	}

	result := db.Save(&dict)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	xsq_net.Success(c)
}

// 修改字典类型
func ChangeDictType(c *gin.Context) {
	var form req.ChangeDictTypeForm

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	var (
		dictType model.DictType
	)

	db := global.DB

	result := db.Where(model.DictType{Code: form.Code}).First(&dictType)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	if result.RowsAffected == 0 {
		xsq_net.ErrorJSON(c, ecode.DataNotExist)
		return
	}

	result = db.Model(&dictType).Updates(form)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	xsq_net.Success(c)

}

// 删除字典类型
func DeleteDictType(c *gin.Context) {

	var form req.DeleteDictTypeForm

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	tx := global.DB.Begin()

	now := time.Now()

	result := tx.Model(&model.DictType{}).Where("code = ?", form.Code).Update("delete_time", &now)

	if result.Error != nil {
		tx.Rollback()
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	result = tx.Model(&model.Dict{}).Where("type_code = ?", form.Code).Update("delete_time", &now)

	if result.Error != nil {
		tx.Rollback()
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	tx.Commit()

	xsq_net.Success(c)
}

// 字典数据列表
func DictList(c *gin.Context) {
	var form req.DictListForm

	err := c.ShouldBind(&form)
	if err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	var (
		dict []model.Dict
	)

	db := global.DB

	result := db.Where(model.Dict{TypeCode: form.Code}).Find(&dict)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	list := make([]*rsp.DictListRsp, 0, result.RowsAffected)
	for _, d := range dict {
		list = append(list, &rsp.DictListRsp{
			TypeCode:   d.TypeCode,
			Code:       d.Code,
			Name:       d.Name,
			Value:      d.Value,
			IsEdit:     d.IsEdit,
			CreateTime: d.CreateTime.Format(timeutil.TimeFormat),
		})
	}

	xsq_net.SucJson(c, list)
}

// 新增字典数据
func CreateDict(c *gin.Context) {
	var form req.CreateDictForm

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	db := global.DB

	var (
		dictType model.DictType
	)

	result := db.Where(&model.DictType{Code: form.TypeCode}).First(&dictType)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	var count int64

	result = db.Model(&model.Dict{}).Where("type_code = ? and code = ?", form.TypeCode, form.Code).Count(&count)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	if count > 0 {
		xsq_net.ErrorJSON(c, ecode.DataAlreadyExist)
		return
	}

	dict := model.Dict{
		Code:     strings.ToLower(form.Code),
		TypeCode: strings.ToLower(form.TypeCode),
		Name:     form.Name,
		Value:    form.Value,
		IsEdit:   form.IsEdit,
	}

	result = db.Create(&dict)
	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	xsq_net.Success(c)
}

// 修改字典数据
func ChangeDict(c *gin.Context) {
	var form req.ChangeDictForm

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	var (
		dict model.Dict
	)

	db := global.DB

	result := db.Where(model.Dict{TypeCode: form.TypeCode, Code: form.Code}).First(&dict)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	if dict.IsEdit == 0 {
		xsq_net.ErrorJSON(c, ecode.DataCannotBeModified)
		return
	}

	result = db.Model(&dict).Updates(form)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	xsq_net.Success(c)

}

// 删除字典数据
func DeleteDict(c *gin.Context) {
	var form req.DeleteDictForm

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	now := time.Now()

	result := global.DB.Model(&model.Dict{}).Where("type_code = ? and code = ?", form.TypeCode, form.Code).Update("delete_time", &now)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	xsq_net.Success(c)
}

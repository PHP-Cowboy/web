package controller

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
	"web/daos"
	"web/forms/req"
	"web/global"
	"web/model"
	"web/utils/ecode"
	"web/utils/xsq_net"
)

// 微信登录
func WxLogin(c *gin.Context) {
	var form req.WxLogin

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	token, err := daos.Wx(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, token)
}

// 登录
func LoginByPwd(c *gin.Context) {
	var form req.LoginByPwd

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	token, err := daos.LoginByPwd(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, gin.H{"token": token})

}

// 验证码登录
func LoginByCode(c *gin.Context) {
	var form req.LoginByCode

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	token, err := daos.LoginByCode(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, gin.H{"token": token})

}

// 注册用户
func Registration(c *gin.Context) {
	var form req.Registration

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	token, err := daos.Registration(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, gin.H{"token": token})
}

// 修改 名称 密码 状态 组织
func ChangeUser(c *gin.Context) {
	var form req.CheckPwdForm

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	var (
		user   model.User
		update model.User
	)

	db := global.DB

	result := db.First(&user, form.Id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			xsq_net.ErrorJSON(c, ecode.UserNotFound)
			return
		}
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	//更新用户密码
	if form.NewPassword != "" {
		update.Password = GenderPwd(form.NewPassword)
	}

	update.Name = form.Name     //更新用户名称
	update.Status = form.Status //更新用户状态

	result = db.Model(&user).Updates(update)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	xsq_net.Success(c)
}

func GenderPwd(pwd string) string {
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(pwd, options)
	return fmt.Sprintf("pbkdf2-sha512$%s$%s", salt, encodedPwd)
}

// 校验登录是否过期
func Check(c *gin.Context) {
	//只需要中间件校验登录是否过期即可
	xsq_net.Success(c)
}

// 建议
func Suggestion(c *gin.Context) {
	var form req.Suggestion

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	uid, ok := c.Get("uid")

	if !ok {
		xsq_net.ErrorJSON(c, ecode.UserNotLogin)
		return
	}

	form.Uid = uid.(int)

	err := daos.Suggestion(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.Success(c)
}

// 支付礼包列表
func PayGiftList(c *gin.Context) {

	list, err := daos.PayGiftList()

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, list)
}

// 生成订单
func Order(c *gin.Context) {
	var form req.Order

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	uid, ok := c.Get("uid")

	if !ok {
		xsq_net.ErrorJSON(c, ecode.UserNotLogin)
		return
	}

	form.Uid = uid.(int)

	orderNo, err := daos.Order(form)

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.SucJson(c, gin.H{"orderNo": orderNo})
}

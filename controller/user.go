package controller

import (
	"context"
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"strconv"
	"strings"
	"time"
	"web/common/constant"
	"web/utils/ecode"
	"web/utils/timeutil"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"

	"web/forms/req"
	"web/forms/rsp"
	"web/global"
	"web/middlewares"
	"web/model"
	"web/utils/xsq_net"
)

// 新增用户
func CreateUser(c *gin.Context) {
	var form req.AddUserForm

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	db := global.DB

	var (
		user model.User
		role model.Role
	)

	result := db.First(&role, form.RoleId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			xsq_net.ErrorJSON(c, ecode.RoleNotFound)
			return
		}
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	user.Password = GenderPwd(form.Password)
	user.Name = form.Name
	user.RoleId = form.RoleId
	user.Role = role.Name

	result = db.Save(&user)
	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	userRsp := rsp.AddUserRsp{
		Id:   user.Id,
		Name: user.Name,
		Role: user.Role,
	}

	xsq_net.SucJson(c, userRsp)
}

// 获取用户列表
func GetUserList(c *gin.Context) {
	var form req.GetUserListForm

	err := c.ShouldBind(&form)
	if err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	var (
		users []model.User
		res   rsp.UserListRsp
	)

	db := global.DB

	result := db.Where("delete_time is null").Find(&users)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	res.Total = result.RowsAffected

	db.Where("delete_time is null").Scopes(model.Paginate(form.Page, form.Size)).Find(&users)

	status := false

	for _, user := range users {
		if user.Status == 1 {
			status = true
		}
		res.List = append(res.List, &rsp.User{
			Id:         user.Id,
			Name:       user.Name,
			Role:       user.Role,
			Status:     status,
			CreateTime: user.CreateTime.Format(timeutil.TimeFormat),
		})
	}

	xsq_net.SucJson(c, res)
}

// 登录
func Login(c *gin.Context) {
	var form req.LoginForm

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	var (
		user model.User
	)

	db := global.DB
	result := db.Where("id = ? and status = 1 and delete_time is null", form.Id).First(&user)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	if result.RowsAffected == 0 {
		xsq_net.ErrorJSON(c, ecode.UserNotFound)
		return
	}

	options := &password.Options{16, 100, 32, sha512.New}

	pwdSlice := strings.Split(user.Password, "$")

	if !password.Verify(form.Password, pwdSlice[1], pwdSlice[2], options) {
		xsq_net.ErrorJSON(c, ecode.PasswordCheckFailed)
		return
	}

	hour := time.Duration(24)

	claims := middlewares.CustomClaims{
		ID:             user.Id,
		Name:           user.Name,
		AuthorityId:    user.RoleId,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(hour * time.Hour).Unix()},
	}

	j := middlewares.NewJwt()
	token, err := j.CreateToken(claims)
	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	//token存入redis
	redisKey := constant.LOGIN_PREFIX + strconv.Itoa(user.Id)
	err = global.Redis.Cli.Set(context.Background(), redisKey, token, hour*60*60*time.Second).Err()
	if err != nil {
		xsq_net.ErrorJSON(c, ecode.RedisFailedToSetData)
		return
	}

	xsq_net.SucJson(c, gin.H{
		"token":  token,
		"roleId": user.RoleId,
		"userId": user.Id,
		"name":   user.Name,
	})
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
	update.RoleId = form.RoleId //更新用户角色

	result = db.Model(&user).Updates(update)

	if result.Error != nil {
		xsq_net.ErrorJSON(c, result.Error)
		return
	}

	xsq_net.Success(c)
}

// 批量删除角色
func BatchDeleteUser(c *gin.Context) {
	var form req.BatchDeleteUserForm

	bindingBody := binding.Default(c.Request.Method, c.ContentType()).(binding.BindingBody)

	if err := c.ShouldBindBodyWith(&form, bindingBody); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	result := global.DB.Model(model.User{}).
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

func GetUser(c *gin.Context) {
	var form req.GetUserForm

	if err := c.ShouldBind(&form); err != nil {
		xsq_net.ErrorJSON(c, ecode.ParamInvalid)
		return
	}

	userModel := model.NewUserModel(global.DB)

	////新增单个
	//user := &model.User{
	//	Password: "123",
	//	Name:     "123",
	//	RoleId:   1,
	//	Role:     "123",
	//	Status:   1,
	//}
	//
	//err := userModel.Save(user)
	//if err != nil {
	//	xsq_net.ErrorJSON(c, err)
	//	return
	//}
	//
	//xsq_net.Success(c)

	//新增多个
	//users := []model.User{
	//	{
	//		Password: "123",
	//		Name:     "123",
	//		RoleId:   1,
	//		Role:     "123",
	//		Status:   1,
	//	},
	//
	//	{
	//		Password: "456",
	//		Name:     "456",
	//		RoleId:   1,
	//		Role:     "456",
	//		Status:   1,
	//	},
	//}
	//
	//err := userModel.CreateInBatches(&users)
	//
	//if err != nil {
	//	xsq_net.ErrorJSON(c, err)
	//	return
	//}
	//
	//xsq_net.Success(c)

	//批量修改
	list := []model.User{
		{
			Base:     model.Base{Id: 10002},
			Password: "789",
			Name:     "789",
			RoleId:   1,
			Role:     "789",
			Status:   1,
		},
	}

	err := userModel.ReplaceCreateInBatches(&list, []string{"name", "role"})

	if err != nil {
		xsq_net.ErrorJSON(c, err)
		return
	}

	xsq_net.Success(c)

	//err, user := userModel.GetFirstByPk(form.Id)
	//
	//if err != nil {
	//	xsq_net.ErrorJSON(c, err)
	//	return
	//}
	//
	//xsq_net.SucJson(c, user)
}

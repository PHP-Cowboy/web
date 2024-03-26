package daos

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"strings"
	"time"
	"web/forms/req"
	"web/forms/rsp"
	"web/global"
	"web/middlewares"
	"web/model"
	"web/utils/ecode"
)

// 密码登录
func LoginByPwd(form req.LoginByPwd) (res *rsp.LoginByPhoneRsp, err error) {
	user := model.User{}

	db := global.DB
	obj := new(model.User)

	err, user = obj.GetUserByPhone(db, form.Phone)

	if err != nil {
		return
	}

	//校验用户状态
	if user.Status != model.UserStatusNormal {
		err = ecode.UserNotFound
		return
	}

	//验证密码
	if !CheckPwd(form.Password) {
		err = ecode.PasswordCheckFailed
		return
	}

	token := ""
	//生成token
	token, err = GetToken(&user)

	res.Id = user.Id
	res.Name = user.Name
	res.Token = token

	return
}

// 注册用户
func Registration(req req.Registration) (res *rsp.LoginByPhoneRsp, err error) {
	user := model.User{}

	db := global.DB
	obj := new(model.User)

	err, user = obj.GetUserByPhone(db, req.Phone)

	//有错误，且不是数据未找到
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	if user.Id > 0 {
		err = ecode.PhoneAlreadyExist
		return
	}

	pwd := GenderPwd(req.Password)

	obj = &model.User{
		Phone:    req.Phone,
		OpenId:   "",
		Name:     req.Phone,
		Password: pwd,
		Status:   model.UserStatusNormal,
	}

	err = obj.Save(db)

	if err != nil {
		return
	}

	token := ""

	token, err = GetToken(obj)
	if err != nil {
		return
	}

	res.Id = user.Id
	res.Name = user.Name
	res.Token = token
	return
}

// 验证密码并生成token
func CheckPwd(pwd string) bool {
	options := &password.Options{16, 100, 32, sha512.New}
	pwdSlice := strings.Split(pwd, "$")
	return password.Verify(pwd, pwdSlice[1], pwdSlice[2], options)
}

// 生成密码
func GenderPwd(pwd string) string {
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(pwd, options)
	return fmt.Sprintf("pbkdf2-sha512$%s$%s", salt, encodedPwd)
}

func GetToken(user *model.User) (token string, err error) {
	hour := time.Duration(24)

	claims := middlewares.CustomClaims{
		ID:             user.Id,
		Name:           user.Name,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(hour * time.Hour).Unix()},
	}

	j := middlewares.NewJwt()
	token, err = j.CreateToken(claims)
	return
}

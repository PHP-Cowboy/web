package daos

import (
	"crypto/sha512"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"io"
	"net/http"
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
func LoginByPwd(form req.LoginByPwd) (token string, err error) {
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

	token, err = GetToken(&user)

	return
}

// 验证码登录
func LoginByCode(form req.LoginByCode) (token string, err error) {
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

	//验证手机号
	err = MsgVerify(form.Phone, form.Code)

	if err != nil {
		return
	}

	//生成token
	token, err = GetToken(&user)

	return
}

// 注册用户
func Registration(req req.Registration) (token string, err error) {
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

	//验证code
	err = MsgVerify(req.Phone, req.Code)

	if err != nil {
		return
	}

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

// 短信验证码校验
func MsgVerify(phone string, code int) (err error) {
	url := "https://webapi.sms.mob.com/sms/verify"

	payload := strings.NewReader(fmt.Sprintf("appkey=%s&phone=%s&zone=86&code=%d", global.ServerConfig.Sms.AppKey, phone, code))

	client := &http.Client{}

	newRequest, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return
	}
	newRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(newRequest)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := &rsp.MsgVerifyRsp{}

	err = json.Unmarshal(body, result)
	if err != nil {
		return
	}

	if result.Status != 200 {
		err = errors.New(result.Error)
	}

	return
}

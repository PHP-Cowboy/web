package daos

import (
	"crypto/sha512"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"web/forms/req"
	"web/forms/rsp"
	"web/global"
	"web/middlewares"
	"web/model"
	"web/utils/ecode"
	"web/utils/timeutil"
)

// 密码登录
func LoginByPwd(form req.LoginByPwd) (token string, err error) {

	data := model.User{}

	db := global.DB
	obj := new(model.User)

	err, data = obj.GetUserByPhone(db, form.Phone)

	if err != nil {
		return
	}

	//校验用户状态
	if data.Status != model.UserStatusNormal {
		err = ecode.UserNotFound
		return
	}

	//验证密码
	if !CheckPwd(form.Password, data.Password) {
		err = ecode.PasswordCheckFailed
		return
	}

	token, err = GetToken(&data)

	return
}

// 验证码登录
func LoginByCode(form req.LoginByCode) (token string, err error) {

	//验证手机号 todo 开放验证码
	//err = MsgVerify(form.Phone, form.Code)

	//if err != nil {
	//	return
	//}

	token, err = Login(req.LoginParams{
		Type:  1,
		Param: form.Phone,
	})

	return
}

func Login(p req.LoginParams) (token string, err error) {

	data := model.User{}

	db := global.DB
	obj := new(model.User)

	var (
		phone   string
		unionId string
	)

	switch p.Type {
	case 1: //手机号
		err, data = obj.GetUserByPhone(db, p.Param)
		phone = p.Param
	case 2: //微信
		err, data = obj.GetUserByUnionId(db, p.Param)
		unionId = p.Param
	}

	if err == nil {
		//校验用户状态
		if data.Status != model.UserStatusNormal {
			err = ecode.UserNotFound
			return
		}
	} else {
		//有错误，错误是数据未找到，走注册逻辑
		if errors.Is(err, gorm.ErrRecordNotFound) {
			data = model.User{
				Phone:    phone,
				UnionId:  unionId,
				Name:     p.Param,
				Password: "",
				Status:   model.UserStatusNormal,
			}

			err = db.Create(&data).Error

			if err != nil {
				return
			}
		} else {
			//有错误且不是数据未找到，直接返回错误
			return
		}
	}

	//生成token
	token, err = GetToken(&data)

	return
}

// 注册用户
func Registration(req req.Registration) (token string, err error) {

	data := model.User{}

	db := global.DB
	obj := new(model.User)

	err, data = obj.GetUserByPhone(db, req.Phone)

	//有错误，且不是数据未找到
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	if data.Id > 0 {
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
		UnionId:  "",
		Name:     req.Phone,
		Password: pwd,
		Status:   model.UserStatusNormal,
	}

	err = obj.Save(db)

	if err != nil {
		return
	}

	token, err = GetToken(obj)
	if err != nil {
		return
	}

	return
}

// 验证密码并生成token
func CheckPwd(pwd, basePwd string) bool {
	options := &password.Options{16, 100, 32, sha512.New}
	pwdSlice := strings.Split(basePwd, "$")
	return password.Verify(pwd, pwdSlice[1], pwdSlice[2], options)
}

// 生成密码
func GenderPwd(pwd string) string {
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(pwd, options)
	return fmt.Sprintf("pbkdf2-sha512$%s$%s", salt, encodedPwd)
}

func GetToken(user *model.User) (token string, err error) {
	if !CheckTime() {
		err = errors.New("已过期")
		return
	}

	claims := middlewares.CustomClaims{
		Id: user.Id,
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
		return
	}

	result := &rsp.MsgVerifyRsp{}

	err = json.Unmarshal(body, result)
	if err != nil {
		return
	}

	if result.Status != 200 {
		return ecode.New(result.Status, result.Error)
	}

	return
}

// 建议
func Suggestion(form req.Suggestion) (err error) {
	db := global.DB

	obj := new(model.Suggestion)

	err = obj.Save(db, form.Uid, form.Msg)
	return
}

// 支付礼包
func PayGiftList() (list []rsp.PayGift, err error) {
	db := global.DB

	obj := new(model.Gift)

	var dataList []model.Gift

	dataList, err = obj.GetList(db)
	if err != nil {
		return
	}

	list = make([]rsp.PayGift, 0, len(dataList))

	for _, d := range dataList {
		list = append(list, rsp.PayGift{
			Id:    d.Id,
			Name:  d.Name,
			Desc:  d.Desc,
			Price: d.Price,
		})
	}
	return
}

// 生成订单
func Order(form req.Order) (orderNo string, err error) {
	db := global.DB

	var gift model.Gift

	giftObj := new(model.Gift)

	gift, err = giftObj.GetFirstById(db, form.Id)

	if err != nil {
		return
	}

	now := time.Now()

	orderNo = now.Format(timeutil.TimeFormNoSplit) + strconv.Itoa(now.Nanosecond())

	obj := &model.Order{
		Uid:     form.Uid,
		OrderNo: orderNo,
		Status:  model.OrderStatusNotPay,
		GiftId:  gift.Id,
		Price:   gift.Price,
		PayType: form.PayType,
	}

	err = obj.Save(db)

	return
}

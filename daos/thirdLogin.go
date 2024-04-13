package daos

import (
	"encoding/json"
	"fmt"
	"web/common/constant"
	"web/forms/req"
	"web/forms/rsp"
	"web/global"
	"web/utils/ecode"
	"web/utils/request"
)

func Wx(form req.Wx) (token string, err error) {
	app := global.ServerConfig.ThirdApp

	url := constant.GetAccessToken + fmt.Sprintf("?appid=%s&secret=%s&code=%s&grant_type=authorization_code", app.AppId, app.Secret, form.Code)

	var b []byte

	b, err = request.Get(url)

	if err != nil {
		return
	}

	var res rsp.WxAccessToken

	err = json.Unmarshal(b, &res)
	if err != nil {
		return
	}

	if res.Errcode != 0 {
		err = ecode.New(res.Errcode, res.Errmsg)
		return
	}

	token, err = Login(req.LoginParams{
		Type:  2,
		Param: res.UnionId,
	})

	return
}

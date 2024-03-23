package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"web/utils/ecode"
	"web/utils/xsq_net"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		xsq_net.ErrorJSON(c, ecode.CreateCaptchaError)
		return
	}
	xsq_net.SucJson(c, gin.H{
		"captchaId": id,
		"path":      b64s,
	})
}

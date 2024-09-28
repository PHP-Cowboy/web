package daos

import (
	"encoding/json"
	"fmt"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"math/rand"
	"strings"
	"sync"
	"time"
	"web/forms/req"
	"web/global"
	"web/model"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
)

var (
	r      *rand.Rand
	once   sync.Once
	client *dysmsapi20170525.Client
)

func init() {
	once.Do(func() {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	})
}

func SendMsg(form req.SendMsg) (err error) {
	code := GetCode()

	obj := new(model.Msg)

	msg := model.Msg{
		Phone: form.Phone,
		Code:  code,
	}

	err = obj.Create(global.DB, &msg)
	if err != nil {
		global.Logger["err"].Errorf("SendMsg obj.Create failed,err:%v", err.Error())
		return
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(form.Phone),
		SignName:      tea.String("验证码短信"),
		TemplateCode:  tea.String("SMS_305495877"),
		TemplateParam: tea.String(fmt.Sprintf("{\"code\":\"%v\"}", code)),
	}

	runtime := &util.RuntimeOptions{}

	client, err = NewClient()

	if err != nil {
		global.Logger["err"].Errorf("sms NewClient failed,err:%v", err.Error())
		return
	}

	_, err = client.SendSmsWithOptions(sendSmsRequest, runtime)

	if err != nil {
		var sdkErr = &tea.SDKError{}
		if _t, ok := err.(*tea.SDKError); ok {
			sdkErr = _t
		} else {
			sdkErr.Message = tea.String(err.Error())
		}
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		// 错误 message
		fmt.Println(tea.StringValue(sdkErr.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(sdkErr.Data)))
		_ = d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, err = util.AssertAsString(sdkErr.Message)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewClient() (result *dysmsapi20170525.Client, err error) {
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考。
	// 建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html。
	aliCfg := global.ServerConfig.AliCloud
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: tea.String(aliCfg.AccessKeyId),
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: tea.String(aliCfg.AccessKeySecret),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Dysmsapi
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")

	return dysmsapi20170525.NewClient(config)
}

func GetCode() string {

	// 生成0到9999之间的随机数
	randomNum := r.Intn(10000)

	// 使用fmt.Sprintf格式化字符串，不足4位时前置填充0
	return fmt.Sprintf("%04d", randomNum)
}

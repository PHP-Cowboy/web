package daos

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"net/url"
	"web/common/constant"
	"web/forms/req"
	"web/global"
	"web/model"
)

func WxPay() {

}

// wap 支付
func ALiTradeWapPay(form req.ALiPay) (payURL string, err error) {
	var (
		client *alipay.Client
		trade  alipay.Trade
		URL    *url.URL
	)

	client, trade, err = ALiPay(form.OrderNo)

	p := alipay.TradeWapPay{
		Trade: trade,
	}

	URL, err = client.TradeWapPay(p)
	if err != nil {
		return
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	payURL = URL.String()

	return
}

// app支付
func ALiTradeAppPay(form req.ALiPay) (result string, err error) {
	var (
		client *alipay.Client
		trade  alipay.Trade
	)

	client, trade, err = ALiPay(form.OrderNo)

	p := alipay.TradeAppPay{
		Trade: trade,
	}

	result, err = client.TradeAppPay(p)
	if err != nil {
		return
	}

	return
}

func ALiPay(orderNo string) (client *alipay.Client, trade alipay.Trade, err error) {
	var (
		privateKey = "MIIEowIBAAKCAQEA1VU8SjAdd8/ey2vjWqVp4CZlD6jKXJzNzByPuO2qn656j/Rlb3HcoirDHNTs9Ar1X3BSPLbH9knpY1sho89FXhsCUTKWA2P5Fti1ZG9b7bvt46c3HtDCk4aYgCa/m2niYy2v0P8/kcfwCvrAccLmwqnQO6dfpzqt+7dCxHmOTg/yE3se/HNbTMl7qQdGOiz6YROs/cSlUhp+FFnkZXxTYluse01I/ruWD2jth7dv+hsXZNFY6ZnchVU/Jp49QtBCFxCLhHGjy/xRz4SsoiJJSA9ouxXehNvHzxIBy9E6qDMtEyMkis1sNfGAjnbreDkZPRJhqMS5LGJvC6fmZsTc9QIDAQABAoIBAC/dH9nFUdc/3Eznt+rukgKomqQYys5coA/bmKN1L4MEMpt5yghnE4mO7pfZbhCTJNp06P6WQ+cOTj7sg2+tYXLrvGlbYC0CkTB/DrOqCIYeVBFvwv42+HEBEdYu226TIyf9aCUMH8clhMXx3Jupjvy9/OplcyCiGccQJvrOCw+YOz92wJm+lsCnlxbnKagerNb25evcDdUrc0NaU+o3b9NgGFQJ2NdlY37vDZljKYm8JwnrLHnEhsAOOkfDY33HlsJ4CVEhPAYHuUvgWQXl7oClw1dm07LTHniEoD6qTKqP5lGyUnHRsB5AQltNRAj85OQrpcuXEU3FbddveEiOxAECgYEA90eSsDmBdkD476MB7ky2eXcrXjmnuSRZa917MLhcyO/j4SY6GGdlPs85S+JF1laafjn5yoQ2t8NNDwtDBh8DQEF73nFVzL7V66hRhoxiTJUrrk1QmmRDHfz9YcV8CBAO4Io5lktxaaUJQThmvBg6NVIgfboGbnfWebPgrXC4qMUCgYEA3NsxsZuueUFT5Olb946bMAg0sDcgGLpEaAhkSS+Saaq9Oz4x4tmqjlTBYW3/Dz8ezxpeRQw6viqvbLq8PPgm53llrBRRX7Mb00mS462V2O2XnGSt87ujOen7riyXTGqEY1ln85KT97P/6ixJBgnnP9NB8XTJfQqBcweZQ69yxnECgYEAx3laMIRDFiS1a4JgdV42uOdT+a+iyCw1YlkJHjUqzAEQH7jZloQZ3UAG5VsDoTK9b+POx/o9taF7UQ5xo8dytOveQ4PDheXP5UkjctleFo8i7Hl78v+8UL2bnLpVIp/pkGKQr9fuBh3WyGD6JvOYSRkWErX0qDpNrmFXdjH/FE0CgYBt9hy5519pM+OHV8DYcwJFWfCRxACRDzy752G3Jp5pSf9Jnd/MqL9Iel6GGfQiLTUzvFcvXbVptWr+YcMTQCJuoXzDvqH9WfoCquEdxvGSvkj8LwyFC+lrDlqnD8CM5YPgy0T8ewB9FfvXJhF9ljSzDIYKqDEhtrdZBHGEx7gHoQKBgEXk/IWyFZ75yExa5nigYmSyrPaeL18TgrKJlCdRYGPlFWJn42C+OUJukNvpTPGBc7VLp57Se4buauuyj9tobS7/EY7MDz7rq3zmh/ZhOn2BC6u2MMi7WyP8Wzwz0rf8UG+Btlhbvm7IiyY3HEimstvZmKRqkPB8etgJICrhq3GQ"
		order      model.Order
	)

	client, err = alipay.New("9021000136623668", privateKey, false)

	if err != nil {
		return
	}

	//查询订单信息
	db := global.DB

	obj := new(model.Order)

	order, err = obj.GetOneByOrderNo(db, orderNo)

	if err != nil {
		return
	}

	trade.NotifyURL = constant.NotifyURL
	trade.ReturnURL = constant.ReturnURL
	trade.ProductCode = "QUICK_MSECURITY_PAY"
	trade.Subject = "伤寒通览-会员开通"
	trade.OutTradeNo = orderNo
	trade.TotalAmount = fmt.Sprintf("%.2f", order.Price)
	return
}

func ALiNotify() {

}

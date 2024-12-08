package daos

import (
	"errors"
	"fmt"
	alipay "github.com/smartwalle/alipay/v3"
	"net/url"
	"strconv"
	"time"
	"web/common/constant"
	"web/forms/req"
	"web/global"
	"web/model"
)

func WxPay() {
	//https://github.com/wechatpay-apiv3/wechatpay-go/blob/main/services/payments/app/api_app.go
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
		privateKey = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCLSDlMLruMO1/7eUWHcn475DmCAf4+3CwSqNU7tbUcfC5QsOJbO3Ki5iq4eyTOLpVhrOavRVbLiGALaFBZCMvdHtgZpNeS68vE9X2IHXN0VFMnAIH/F5DMgXYegbScJtQWSqqBYXKjDVn3u/WtKjKKnENMb5hRl7RXq8gzXgJooooLi1FJZFxQ0M0tSfKOYFeU2lFx5BPTKvW3B9IDmY3DZdIz8/VNtKn60/KYkx91wYev60M3fxni4ClqXBGR5QgRT7fl0Z7kfyuxLLPje/AeDk1NZZslFFsOvp8wEbz9KXseQkmeRHXoiuQM6vznNHB4x1SNBmSmLWCIknEVavVnAgMBAAECggEAEEvfCF7fasT3sgC6deKbin5ljCSXjttL/NLsQBQ+oJqXALHxiiMmgpDTpsohwE4LBGaqhw9h3GaUdkE6RXCWCcU2G6oB3xrfuKfQjECF5bKIGCQjEam8M3FiVKdwbmTMo98QnBq+wv0o7ad8U0qAp18pMKRN+h3559gdkIULWxqdJf+Jg2gNBHKmc71EDiyN06ysD44fgnFeKhiJww0bXEwhkM4RtEps5T/C4Nca+O7lgYsfHCx6CNBnULL15nb3yb9EdnCc55vYS2YFVywM9ynSqotr3/1jzeGbeMrwB+g6uhPcqd75DWQqdNJejNIBeOQLn6ngu+7dh6LplquNgQKBgQDLJ4f6rsyszNWk4PgBAOolpOSphgUhLiQINC++Ny2uEjaMQ9ZIjXWp8sYRAiTzn9Nxq4jlo/PKwfwIpuwrYSdNtAAmvyl8ixRU8DMc0x1eSbqYeu2CNF64gsaMKWA5JIqnBkoaxDbV6QkD8xNwHQkyPillp4UOZk4vwp8/IdwVUQKBgQCvg05Fi7eWhsvrlxzvKuMqUVygKxjNgEYSp1gmG4QZEXxDxSCpmcgbr35z7fe5yh2t1vWByi04yKwJOSUplN5e3UaLwJhDujXTcMXASeErS57UU2an+8BTP0TQ171fbYLUjP/WQW1G6dU173tLHTamshYakw9HXapCUTgyCDgRNwKBgAgy9zz50w7iUXkPD1NTCyAr0IK7Avf82b1vBcJCjdhpqcPtdPeGpgoO30mOcZlXiPwcdTD3df5n59kdi4iQqjprmln1+yGop1BqRVXTe2rzxNEknPNVU8FUndjITrbfr6zgBxufcasedZuwMseilF1vvtH4edlD7cHls/GjVzORAoGBAKd53A5wq9Zu+dbscyek/O88g+4yMZCmfjb+iBaU2xAwpyrjx68OluwUIc3tnR0jx/5ByDL5AP3WPpLufI+ZBjfvXYmtVZ4K6aHBUnGSDdKKbhydGK/igjrtfHTr6EBO9zBq+xUviV7JcjHtOvEKB0Q5Wu4wtbubnQN7DOxw+08xAoGBAJ5zF/nx1Lv1nEMCdp2jQVUYJJE4pq3ZybN09StTO6tietBbihbHRjJLeEg4Fdl/w7ZEkz4bohlrrVIPz4NEASqLw3skmjx4YBBmFff/xRlT3iK2LLLwBHaLzaW5ToJdvfn2NKcp8amSYkxniXrA1I2fE4Gg0+qFBgmxd8LUR3x6"
		//privateKey = "MIIEowIBAAKCAQEA1VU8SjAdd8/ey2vjWqVp4CZlD6jKXJzNzByPuO2qn656j/Rlb3HcoirDHNTs9Ar1X3BSPLbH9knpY1sho89FXhsCUTKWA2P5Fti1ZG9b7bvt46c3HtDCk4aYgCa/m2niYy2v0P8/kcfwCvrAccLmwqnQO6dfpzqt+7dCxHmOTg/yE3se/HNbTMl7qQdGOiz6YROs/cSlUhp+FFnkZXxTYluse01I/ruWD2jth7dv+hsXZNFY6ZnchVU/Jp49QtBCFxCLhHGjy/xRz4SsoiJJSA9ouxXehNvHzxIBy9E6qDMtEyMkis1sNfGAjnbreDkZPRJhqMS5LGJvC6fmZsTc9QIDAQABAoIBAC/dH9nFUdc/3Eznt+rukgKomqQYys5coA/bmKN1L4MEMpt5yghnE4mO7pfZbhCTJNp06P6WQ+cOTj7sg2+tYXLrvGlbYC0CkTB/DrOqCIYeVBFvwv42+HEBEdYu226TIyf9aCUMH8clhMXx3Jupjvy9/OplcyCiGccQJvrOCw+YOz92wJm+lsCnlxbnKagerNb25evcDdUrc0NaU+o3b9NgGFQJ2NdlY37vDZljKYm8JwnrLHnEhsAOOkfDY33HlsJ4CVEhPAYHuUvgWQXl7oClw1dm07LTHniEoD6qTKqP5lGyUnHRsB5AQltNRAj85OQrpcuXEU3FbddveEiOxAECgYEA90eSsDmBdkD476MB7ky2eXcrXjmnuSRZa917MLhcyO/j4SY6GGdlPs85S+JF1laafjn5yoQ2t8NNDwtDBh8DQEF73nFVzL7V66hRhoxiTJUrrk1QmmRDHfz9YcV8CBAO4Io5lktxaaUJQThmvBg6NVIgfboGbnfWebPgrXC4qMUCgYEA3NsxsZuueUFT5Olb946bMAg0sDcgGLpEaAhkSS+Saaq9Oz4x4tmqjlTBYW3/Dz8ezxpeRQw6viqvbLq8PPgm53llrBRRX7Mb00mS462V2O2XnGSt87ujOen7riyXTGqEY1ln85KT97P/6ixJBgnnP9NB8XTJfQqBcweZQ69yxnECgYEAx3laMIRDFiS1a4JgdV42uOdT+a+iyCw1YlkJHjUqzAEQH7jZloQZ3UAG5VsDoTK9b+POx/o9taF7UQ5xo8dytOveQ4PDheXP5UkjctleFo8i7Hl78v+8UL2bnLpVIp/pkGKQr9fuBh3WyGD6JvOYSRkWErX0qDpNrmFXdjH/FE0CgYBt9hy5519pM+OHV8DYcwJFWfCRxACRDzy752G3Jp5pSf9Jnd/MqL9Iel6GGfQiLTUzvFcvXbVptWr+YcMTQCJuoXzDvqH9WfoCquEdxvGSvkj8LwyFC+lrDlqnD8CM5YPgy0T8ewB9FfvXJhF9ljSzDIYKqDEhtrdZBHGEx7gHoQKBgEXk/IWyFZ75yExa5nigYmSyrPaeL18TgrKJlCdRYGPlFWJn42C+OUJukNvpTPGBc7VLp57Se4buauuyj9tobS7/EY7MDz7rq3zmh/ZhOn2BC6u2MMi7WyP8Wzwz0rf8UG+Btlhbvm7IiyY3HEimstvZmKRqkPB8etgJICrhq3GQ"
		order model.Order
	)

	client, err = alipay.New("2021004132603404", privateKey, true)

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

func Notify(orderNo string) (err error) {
	db := global.DB

	orderObj := new(model.Order)

	orderData, err := orderObj.GetOneByOrderNo(db, orderNo)
	if err != nil {
		global.Logger["err"].Errorf("ALiNotify orderObj.GetOneByOrderNo failed,err:%s,order_no:%s", err.Error(), orderNo)
		return
	}

	giftObj := new(model.Gift)

	giftData, err := giftObj.GetFirstById(db, orderData.GiftId)
	if err != nil {
		global.Logger["err"].Errorf("ALiNotify giftObj.GetFirstById failed,err:%s,gift id:%v", err.Error(), orderData.GiftId)
		return err
	}

	userObj := new(model.User)

	userData, err := userObj.GetFirstByPk(db, orderData.Uid)
	if err != nil {
		global.Logger["err"].Errorf("ALiNotify userObj.GetFirstByPk failed,err:%s,user id:%v", err.Error(), orderData.Uid)
		return
	}

	//永久会员
	if giftData.MemberMonth == 0 {
		userData.MemberLevel = model.UserMemberLevelPermanent
	} else if giftData.MemberMonth > 0 {
		userData.MemberLevel = model.UserMemberLevelRegular

		// 获取当前时间
		now := time.Now()
		// 计算下一天凌晨零点的时间
		nextDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, 1)
		// 增加时间月份数
		targetTime := nextDay.AddDate(0, giftData.MemberMonth, 0)
		// 获取时间戳（秒）
		userData.MemberExpire = int(targetTime.Unix())
	} else {
		err = errors.New("礼包数据异常:gift id:" + strconv.Itoa(giftData.Id))
		return
	}

	err = userData.Save(db)
	if err != nil {
		global.Logger["err"].Errorf("ALiNotify userData.Save failed,err:%s", err.Error())
		return
	}

	return
}

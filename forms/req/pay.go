package req

type ALiPay struct {
	OrderNo string `json:"order_no" binding:"required"`
}

type ALiNotify struct {
	OrderNo string `json:"order_no" binding:"required"`
}
type WxNotify struct {
	OrderNo string `json:"order_no" binding:"required"`
}

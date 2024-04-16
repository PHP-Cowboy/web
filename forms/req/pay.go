package req

type ALiPay struct {
	OrderNo string `json:"order_no" binding:"required"`
}

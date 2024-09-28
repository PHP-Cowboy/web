package req

type SendMsg struct {
	Phone string `json:"phone" binding:"required"`
}

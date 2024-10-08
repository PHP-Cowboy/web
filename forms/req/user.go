package req

type AddUserForm struct {
	WarehouseId int    `forms:"warehouse_id" json:"warehouse_id" binding:"required"`
	RoleId      int    `forms:"role_id" json:"role_id" binding:"required"`
	Name        string `forms:"name" json:"name" binding:"required"`
	Password    string `forms:"password" json:"password" binding:"required"`
}

type GetUserListForm struct {
	Paging
	WarehouseId int `forms:"warehouse_id" json:"warehouse_id"`
}

type LoginByPwd struct {
	CaptchaId string `json:"captcha_id" binding:"required"`
	Captcha   string `json:"captcha" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type LoginByCode struct {
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

type LoginParams struct {
	Type  int    `json:"type"`
	Param string `json:"param"`
}

type Registration struct {
	Phone    string `json:"phone" binding:"required"`
	Code     int    `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CheckPwdForm struct {
	Id          int    `json:"id" binding:"required"`
	NewPassword string `json:"new_password"`
	Name        string `json:"name"`
	Status      int    `json:"status"`
}

type BatchDeleteUserForm struct {
	Ids []int `json:"ids" binding:"required"`
}

type WarehouseUserCountForm struct {
}

type GetPickerListReq struct {
	RoleId int `json:"role_id"`
}

type GetUserForm struct {
	Id int `json:"id" form:"id" binding:"required"`
}

type Suggestion struct {
	Uid int
	Msg string `json:"msg" binding:"required"`
}

type PayGiftList struct {
}

type Order struct {
	Id      int `json:"id" binding:"required"`
	PayType int `json:"pay_type" binding:"required"`
	Uid     int
}

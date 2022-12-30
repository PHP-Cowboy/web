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

type LoginForm struct {
	Id       string `forms:"id" json:"id" binding:"required"`
	Password string `json:"password" forms:"password" binding:"required"`
}

type CheckPwdForm struct {
	Id          int    `json:"id" binding:"required"`
	NewPassword string `json:"new_password"`
	Name        string `json:"name"`
	Status      int    `json:"status"`
	RoleId      int    `json:"role_id"`
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

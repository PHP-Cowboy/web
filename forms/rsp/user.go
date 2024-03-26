package rsp

type AddUserRsp struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	WarehouseId int    `json:"warehouse_id"`
	Status      int    `json:"status"`
	CreateTime  string `json:"create_time"`
}

type UserListRsp struct {
	Total int64   `json:"total"`
	List  []*User `json:"list"`
}

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	WarehouseId int    `json:"warehouse_id"`
	Status      bool   `json:"status"` //bool
	CreateTime  string `json:"create_time"`
}

type LoginByPhoneRsp struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
	Name  string `json:"name"`
}

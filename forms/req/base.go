package req

type Paging struct {
	Page int `form:"page" json:"page" binding:"required,gt=0"`
	Size int `form:"size" json:"size" binding:"required,gt=0,lte=500"`
}

type Id struct {
	Id int `json:"id" form:"id" binding:"required,gt=0"`
}

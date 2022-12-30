package req

type Paging struct {
	Page int `forms:"page" json:"page" binding:"required,gt=0"`
	Size int `forms:"size" json:"size" binding:"required,gt=0,lte=500"`
}

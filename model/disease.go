package model

import "gorm.io/gorm"

// 疾病
type Disease struct {
	Base
	CategoryId int    `gorm:"type:int(11);not null;default:0;comment:分类id"`
	Title      string `gorm:"type:varchar(32);not null;comment:标题"`
	Content    string `gorm:"type:text;not null;comment:内容"`
	Sort       int    `gorm:"type:int(11);not null;default:0;comment:排序"`
}

func (t *Disease) GetListByCategoryId(db *gorm.DB, categoryId int) (list []Disease, err error) {
	err = db.Model(t).Where("category_id = ?", categoryId).Order("sort asc").Find(&list).Error
	return
}

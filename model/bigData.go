package model

import (
	"gorm.io/gorm"
)

// 大数据
type BigData struct {
	Base
	CategoryId int    `gorm:"type:int(11);not null;comment:分类id"`
	Name       string `gorm:"type:varchar(32);not null;default:'';comment:名称"`
	Num        int    `gorm:"type:int(11);not null;comment:数量"`
}

func (t *BigData) GetListByCategoryId(db *gorm.DB, categoryId int) (list []BigData, err error) {
	err = db.Model(t).Where("category_id = ?", categoryId).Find(&list).Error

	return
}

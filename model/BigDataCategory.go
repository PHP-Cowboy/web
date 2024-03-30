package model

import "gorm.io/gorm"

// 大数据分类
type BigDataCategory struct {
	Base
	Name string `gorm:"type:varchar(32);not null;default:'';comment:名称"`
}

func (t *BigDataCategory) GetOneById(db *gorm.DB, id int) (category BigDataCategory, err error) {
	err = db.Model(t).First(&category, id).Error
	return
}

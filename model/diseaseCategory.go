package model

import (
	"gorm.io/gorm"
)

// 疾病分类
type DiseaseCategory struct {
	Base
	ParentId int    `gorm:"type:int(11);not null;default:0;comment:父级分类id"`
	Name     string `gorm:"type:varchar(32);not null;comment:分类名称"`
}

func (t *DiseaseCategory) GetList(db *gorm.DB) (list []DiseaseCategory, err error) {
	err = db.Model(t).Find(&list).Error
	return
}

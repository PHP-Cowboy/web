package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 典籍分类
type ClassicsCategory struct {
	Base
	Name    string `gorm:"type:varchar(32);not null;default:'';comment:分类名称"`
	Picture string `gorm:"type:varchar(255);not null;default:'';comment:分类图片地址"`
}

func (t *ClassicsCategory) GetPageList(db *gorm.DB, form req.ClassicsCategoryList) (total int64, list []ClassicsCategory, err error) {
	localDb := db.Model(t).Where(&ClassicsCategory{Name: form.Name})

	err = localDb.Count(&total).Error

	if err != nil {
		return
	}

	err = localDb.Scopes(Paginate(form.Page, form.Size)).Find(&list).Error

	return
}

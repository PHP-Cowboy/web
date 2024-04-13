package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 常用方剂分类
type CommonlyPrescriptionCategory struct {
	Base
	Name    string `gorm:"type:varchar(32);not null;comment:分类名称"`
	Picture string `gorm:"type:varchar(255);not null;default:'';comment:图片"` //图片
}

func (t *CommonlyPrescriptionCategory) GetPageList(db *gorm.DB, form req.CommonlyPrescriptionCategoryList) (total int64, list []CommonlyPrescriptionCategory, err error) {
	local := db.Model(t)

	if form.KeyWords != "" {
		local.Where("name like ?", "%"+form.KeyWords+"%")
	}

	err = local.Count(&total).Error

	if err != nil {
		return
	}

	err = db.Model(t).Scopes(Paginate(form.Page, form.Size)).Find(&list).Error
	return
}

package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 名医心法
type MindMethod struct {
	Base
	Title   string `gorm:"type:varchar(32);not null;comment:标题"`
	Doctor  string `gorm:"type:varchar(32);not null;comment:医生"`
	Content string `gorm:"type:text;not null;comment:内容"`
}

func (t *MindMethod) GetFirst(db *gorm.DB, form req.Id) (m MindMethod, err error) {
	err = db.Model(t).First(&m, form.Id).Error
	return
}

func (t *MindMethod) GetPageList(db *gorm.DB, form req.MindMethodList) (total int64, list []MindMethod, err error) {
	local := db.Model(t)

	if form.KeyWords != "" {
		local.Where("title like ?", "%"+form.KeyWords+"%").
			Or("doctor like ?", "%"+form.KeyWords+"%").
			Or("content like ?", "%"+form.KeyWords+"%")
	}

	err = local.Count(&total).Error

	if err != nil {
		return
	}

	err = db.Model(t).Scopes(Paginate(form.Page, form.Size)).Find(&list).Error
	return
}

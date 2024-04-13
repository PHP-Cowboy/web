package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 思维导图
type MindMap struct {
	Base
	Title   string `gorm:"type:varchar(32);not null;comment:标题"`
	Picture string `gorm:"type:varchar(255);not null;comment:图片地址"`
}

func (t *MindMap) GetPageList(db *gorm.DB, form req.MindMapList) (total int64, list []MindMap, err error) {
	localDb := db.Model(t).Where("title like ?", "%"+form.KeyWords+"%")

	err = localDb.Count(&total).Error

	if err != nil {
		return
	}

	err = localDb.Scopes(Paginate(form.Page, form.Size)).Find(&list).Error
	return
}

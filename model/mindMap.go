package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

type MindMap struct {
	Base
	Title   string `gorm:"type:varchar(32);not null;comment:标题"`
	Picture string `gorm:"type:varchar(255);not null;comment:图片地址"`
}

func (t *MindMap) TableName() string {
	return "mind_map"
}

func (t *MindMap) GetPageList(db *gorm.DB, form req.MindMapList) (total int64, list []MindMap, err error) {
	err = db.Model(t).Where("title like %?%", form.KeyWords).Scopes(Paginate(form.Page, form.Size)).Find(&list).Error
	return
}

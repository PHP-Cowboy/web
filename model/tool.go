package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 工具
type Tool struct {
	Base
	Name    string `gorm:"type:varchar(32);not null;default:'';comment:名称"`
	Desc    string `gorm:"type:varchar(32);not null;default:'';comment:描述"`
	Picture string `gorm:"type:varchar(32);not null;default:'';comment:描述"`
	Router  string `gorm:"type:varchar(32);not null;default:'';comment:跳转路由"`
}

func (t *Tool) GetPageList(db *gorm.DB, form req.ToolList) (total int64, list []Tool, err error) {
	localDb := db.Model(t)

	err = localDb.Count(&total).Error

	if err != nil {
		return
	}

	err = localDb.Scopes(Paginate(form.Page, form.Size)).Find(&list).Error

	return
}

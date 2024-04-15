package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 试题类别
type QuestionCategory struct {
	Base
	Name string `gorm:"type:varchar(32);not null;default:'';comment:名称"`
}

func (t *QuestionCategory) GetPageList(db *gorm.DB, form req.QuestionCategoryList) (total int64, list []QuestionCategory, err error) {
	local := db.Model(t)

	err = local.Count(&total).Error

	if err != nil {
		return
	}

	err = db.Model(t).Scopes(Paginate(form.Page, form.Size)).Find(&list).Error
	return
}

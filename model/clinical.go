package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 临床
type Clinical struct {
	Base
	Symptom string `gorm:"type:varchar(32);not null;default:'';comment:症状"`
}

func (t *Clinical) GetPageList(db *gorm.DB, form req.ClinicalList) (total int64, list []Clinical, err error) {
	localDb := db.Model(t)

	err = localDb.Count(&total).Error

	if err != nil {
		return
	}

	err = localDb.Scopes(Paginate(form.Page, form.Size)).Find(&list).Error

	return
}

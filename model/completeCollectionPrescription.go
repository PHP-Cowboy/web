package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 方剂大全
type CompleteCollectionPrescription struct {
	Base
	Title      string `gorm:"type:varchar(32);not null;comment:方剂标题"`
	Provenance string `gorm:"type:varchar(32);not null;default:'';comment:方剂出处"` //出处
	Content    string `gorm:"type:varchar(255);not null;default:'';comment:内容"`  //内容
}

func (t *CompleteCollectionPrescription) GetFirst(db *gorm.DB, form req.Id) (m CompleteCollectionPrescription, err error) {
	err = db.Model(t).First(&m, form.Id).Error
	return
}

func (t *CompleteCollectionPrescription) GetPageList(db *gorm.DB, form req.CompleteCollectionPrescriptionList) (total int64, list []CompleteCollectionPrescription, err error) {
	local := db.Model(t)

	if form.KeyWords != "" {
		local.Where("title like ?", "%"+form.KeyWords+"%").
			Or("provenance like ?", "%"+form.KeyWords+"%").
			Or("content like ?", "%"+form.KeyWords+"%")
	}

	err = local.Count(&total).Error

	if err != nil {
		return
	}

	err = db.Model(t).Scopes(Paginate(form.Page, form.Size)).Find(&list).Error
	return
}

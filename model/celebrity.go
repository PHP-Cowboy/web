package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 名家
type Celebrity struct {
	Base
	NotableDoctor  string `gorm:"type:varchar(16);not null;default:'';comment:医家"`
	Provenance     string `gorm:"type:varchar(32);not null;default:'';comment:出处"`
	Content        string `gorm:"type:varchar(255);not null;default:'';comment:内容"`
	PrescriptionId int    `gorm:"type:int(11);not null;comment:所属方剂id"`
}

func (t *Celebrity) GetFirst(db *gorm.DB, form req.Id) (m Celebrity, err error) {
	err = db.Model(t).First(&m, form.Id).Error
	return
}

func (t *Celebrity) GetFirstByPrescriptionId(db *gorm.DB, prescriptionId int) (m Celebrity, err error) {
	err = db.Model(t).Where("prescription_id = ?", prescriptionId).First(&m).Error
	return
}

func (t *Celebrity) GetPageList(db *gorm.DB, form req.PrescriptionCelebrityList) (total int64, list []Celebrity, err error) {
	local := db.Model(t).Where(Celebrity{PrescriptionId: form.PrescriptionId})

	if form.KeyWords != "" {
		local.Where("notable_doctor like ?", "%"+form.KeyWords+"%").
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

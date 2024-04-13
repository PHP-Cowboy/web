package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 方剂
type Prescription struct {
	Base
	Name          string `gorm:"type:varchar(64);not null;default:'';comment:名称"`
	Symptom       string `gorm:"type:varchar(64);not null;default:'';comment:症状"`
	TongueQuality string `gorm:"type:varchar(64);not null;default:'';comment:舌质"`
	CoatedTongue  string `gorm:"type:varchar(64);not null;default:'';comment:舌苔"`
	Pulse         string `gorm:"type:varchar(64);not null;default:'';comment:脉象"`
	ModernDisease string `gorm:"type:varchar(64);not null;default:'';comment:现代疾病"`
	SymptomPic    string `gorm:"type:varchar(255);not null;default:'';comment:症状图"`
	PulsePic      string `gorm:"type:varchar(64);not null;default:'';comment:脉象图"`
	TonguePic     string `gorm:"type:varchar(64);not null;default:'';comment:舌象图"`
}

func (t *Prescription) GetFirst(db *gorm.DB, form req.Id) (m Prescription, err error) {
	err = db.Model(t).First(&m, form.Id).Error
	return
}

func (t *Prescription) GetPageList(db *gorm.DB, form req.PrescriptionList) (total int64, list []Prescription, err error) {
	local := db.Model(t)

	if form.KeyWords != "" {
		local.Where("name like ?", "%"+form.KeyWords+"%").
			Or("symptom like ?", "%"+form.KeyWords+"%").
			Or("tongue_quality like ?", "%"+form.KeyWords+"%").
			Or("coated_tongue like ?", "%"+form.KeyWords+"%").
			Or("pulse like ?", "%"+form.KeyWords+"%").
			Or("modern_disease like ?", "%"+form.KeyWords+"%")
	}

	err = local.Count(&total).Error

	if err != nil {
		return
	}

	err = db.Model(t).Scopes(Paginate(form.Page, form.Size)).Find(&list).Error
	return
}

func (t *Prescription) GetRandList(db *gorm.DB, form req.PrescriptionRandList) (list []Prescription, err error) {
	err = db.Raw("select * from prescription order by RAND() limit ?", form.Size).Scan(&list).Error
	return
}

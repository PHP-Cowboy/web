package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 医疗案例
type MedicalCases struct {
	Base
	Symptom      string `gorm:"type:varchar(255);not null;default:'';comment:症状"` //症状
	Prescription string `gorm:"type:varchar(255);not null;default:'';comment:方药"` //方药
	Content      string `gorm:"type:varchar(255);not null;default:'';comment:内容"` //内容
	Provenance   string `gorm:"type:varchar(32);not null;default:'';comment:出处"`  //出处
}

func (t *MedicalCases) Save(db *gorm.DB) (err error) {
	err = db.Model(t).Save(t).Error
	return
}

func (t *MedicalCases) UpdateById(db *gorm.DB, id uint64, mp map[string]interface{}) (err error) {
	err = db.Model(t).Where("id = ?", id).Updates(mp).Error
	return
}

func (t *MedicalCases) DeleteById(db *gorm.DB, id uint64) (err error) {
	err = db.Model(t).Delete("id = ?", id).Error
	return
}

func (t *MedicalCases) Count(db *gorm.DB) (total int64, err error) {
	err = db.Model(t).Where(t).Count(&total).Error
	return
}

func (t *MedicalCases) GetFirstById(db *gorm.DB, id int) (MedicalCases MedicalCases, err error) {
	err = db.Model(t).Where("id = ?", id).First(&MedicalCases).Error
	return
}

func (t *MedicalCases) GetList(db *gorm.DB) (list []MedicalCases, err error) {
	err = db.Model(t).Where(t).Find(&list).Error
	return
}

func (t *MedicalCases) GetPageList(db *gorm.DB, form req.MedicalCasesList) (list []MedicalCases, err error) {
	err = db.Model(t).Where(t).Scopes(Paginate(form.Page, form.Size)).Find(&list).Error
	return
}

func (t *MedicalCases) GetLikePageList(db *gorm.DB, form req.MedicalCasesList) (list []MedicalCases, err error) {
	localDb := db.Model(t)
	if form.KeyWords != "" {
		localDb.Where("symptom like '%" + form.KeyWords + "%'" + " or prescription like '%" + form.KeyWords + "%'" + " or provenance like '%" + form.KeyWords + "%'")
	}

	err = localDb.Scopes(Paginate(form.Page, form.Size)).Find(&list).Error
	return
}

func (t *MedicalCases) GetLikeCount(db *gorm.DB, form req.MedicalCasesList) (total int64, err error) {
	localDb := db.Model(t)
	if form.KeyWords != "" {
		localDb.Where("symptom like '%" + form.KeyWords + "%'" + " or prescription like '%" + form.KeyWords + "%'" + " or provenance like '%" + form.KeyWords + "%'")
	}

	err = localDb.Count(&total).Error
	return
}

package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 方剂 prescription
type Formula struct {
	Base
	Name         string `gorm:"type:varchar(255);not null;default:'';comment:方剂名称"`
	Abbreviation string `gorm:"type:varchar(255);not null;default:'';comment:方剂名称缩写"`
	Content      string `gorm:"type:varchar(255);not null;default:'';comment:内容"`
	UserId       int    `gorm:"type:int(11);not null;default:0;comment:用户id"` //方剂所属用户，为0时代表为系统自有方剂，其他为用户方剂
}

func (t *Formula) Create(db *gorm.DB, data Formula) error {
	return db.Model(t).Create(&data).Error
}

func (t *Formula) GetOne(db *gorm.DB, form req.Id) (data Formula, err error) {
	err = db.Model(t).First(&data, form.Id).Error
	return
}

func (t *Formula) GetPageList(db *gorm.DB, form req.FormulaList) (total int64, dataList []Formula, err error) {
	localDb := db.Model(t)

	if form.IsMy {
		localDb.Where("user_id = ?", form.UserId)
	}

	if form.Name != "" {
		localDb.Where("name like ?", "%"+form.Name+"%")
		localDb.Where("abbreviation like ?", "%"+form.Name+"%")
	}

	err = localDb.Count(&total).Error

	if err != nil {
		return
	}

	err = localDb.Scopes(Paginate(form.Page, form.Size)).Find(&dataList).Error

	if err != nil {
		return
	}

	return
}

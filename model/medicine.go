package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 中药
type Medicine struct {
	Base
	Name   string `gorm:"type:varchar(255);not null;default:'';comment:中药名称"`
	UserId int    `gorm:"type:int(11);not null;default:0;comment:用户id"` //中药所属用户，为0时代表为系统自有方剂，其他为用户方剂
}

func (t *Medicine) Create(db *gorm.DB, data Medicine) error {
	return db.Model(t).Create(&data).Error
}

func (t *Medicine) GetPageList(db *gorm.DB, form req.MedicineList) (total int64, dataList []Medicine, err error) {
	localDb := db.Model(t).Where("user_id = ? or user_id = 0", form.UserId)

	if form.Name != "" {
		localDb.Where("name like ?", "%"+form.Name+"%")
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

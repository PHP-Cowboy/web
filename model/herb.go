package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 药物表
type Herb struct {
	Base
	Name         string `gorm:"type:varchar(255);not null;default:'';comment:药物名称"`
	Abbreviation string `gorm:"type:varchar(255);not null;default:'';comment:药物名称缩写"`
	Nature       string `gorm:"type:varchar(255);not null;default:'';comment:属性 "`
	Brief        string `gorm:"type:varchar(255);not null;default:'';comment:功效"`
	UserId       int    `gorm:"type:int(11);not null;default:0;comment:用户id"` //中药所属用户，为0时代表为系统自有方剂，其他为用户方剂
}

func (t *Herb) Create(db *gorm.DB, data Herb) error {
	return db.Model(t).Create(&data).Error
}

func (t *Herb) GetPageList(db *gorm.DB, form req.HerbList) (total int64, dataList []Herb, err error) {
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

func (t *Herb) GetListByIds(db *gorm.DB, ids []int) (dataList []Herb, err error) {
	err = db.Model(t).Where("id in (?)", ids).Find(&dataList).Error

	return
}

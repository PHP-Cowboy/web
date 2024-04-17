package model

import "gorm.io/gorm"

// 礼包
type Gift struct {
	Base
	Name  string  `gorm:"type:varchar(64);not null;default:'';comment:名称"`
	Desc  string  `gorm:"type:varchar(32);not null;default:'';comment:描述"`
	Price float64 `gorm:"type:decimal(10,2);comment:价格"`
}

func (t *Gift) GetFirstById(db *gorm.DB, id int) (data Gift, err error) {
	err = db.Model(t).Where("id = ?", id).First(&data).Error
	return
}

func (t *Gift) GetList(db *gorm.DB) (list []Gift, err error) {
	err = db.Model(t).Where(t).Find(&list).Error
	return
}

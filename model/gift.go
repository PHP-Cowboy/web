package model

import "gorm.io/gorm"

// 礼包
type Gift struct {
	Base
	Name        string  `gorm:"type:varchar(64);not null;default:'';comment:名称"`
	Desc        string  `gorm:"type:varchar(32);not null;default:'';comment:描述"`
	Price       float64 `gorm:"type:decimal(10,2);comment:价格"`
	MemberMonth int     `gorm:"type:int;not null;comment:会员月数:0:永久,其他数字表示N个月"`
}

func (t *Gift) GetFirstById(db *gorm.DB, id int) (data Gift, err error) {
	err = db.Model(t).Where("id = ?", id).First(&data).Error
	return
}

func (t *Gift) GetList(db *gorm.DB) (list []Gift, err error) {
	err = db.Model(t).Where(t).Find(&list).Error
	return
}

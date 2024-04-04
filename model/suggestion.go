package model

import "gorm.io/gorm"

// 建议
type Suggestion struct {
	Base
	Uid int    `gorm:"type:int(11);not null;comment:用户id"`
	Msg string `gorm:"type:varchar(255);not null;comment:建议内容"`
}

func (t *Suggestion) Save(db *gorm.DB, uid int, msg string) (err error) {
	err = db.Model(t).Save(&Suggestion{Uid: uid, Msg: msg}).Error
	return
}

package model

import "gorm.io/gorm"

type Msg struct {
	Base
	Phone  string `gorm:"type:varchar(16);index;not null;comment:手机号"`
	Code   string `gorm:"type:char(4);not null;comment:验证码"`
	Status int    `gorm:"type:tinyint(1);not null;default:1;comment:状态:1:未使用,2:已使用"`
}

const (
	MsgStatus            = iota
	MsgStatusNotUsed     //未使用
	MsgStatusUsedAlready //已使用
)

func (t *Msg) Create(db *gorm.DB, data *Msg) (err error) {
	err = db.Model(t).Create(&data).Error
	return
}

func (t *Msg) UpdateStatusById(db *gorm.DB, id int) error {
	return db.Model(t).Where("id = ?", id).Update("status", MsgStatusUsedAlready).Error
}

func (t *Msg) GetCodeByPhone(db *gorm.DB, phone string) (data Msg, err error) {
	err = db.Model(t).Where("phone = ? and status = ?", phone, MsgStatusNotUsed).Order("id desc").First(&data).Error
	return
}

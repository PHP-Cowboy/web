package model

import (
	"errors"
	"gorm.io/gorm"
	"web/utils/ecode"
)

// 用户
type User struct {
	Base
	Phone    string `gorm:"type:varchar(16);unique;not null;comment:手机号"`
	OpenId   string `gorm:"type:varchar(64);unique;not null;comment:微信openid"`
	Name     string `gorm:"type:varchar(16);not null;comment:姓名"`
	Password string `gorm:"type:varchar(100);not null;comment:密码"`
	Status   int    `gorm:"type:tinyint;not null;default:1;comment:状态:0:未知,1:正常,2:禁用"`
}

const (
	UserStatus          = iota
	UserStatusNormal    //正常
	UserStatusForbidden //禁用
)

// 保存用户
func (t *User) Save(db *gorm.DB, user *User) (err error) {
	err = db.Model(t).Save(user).Error
	return
}

// 批量创建用户
func (t *User) CreateInBatches(db gorm.DB, list *[]User) (err error) {
	err = db.Model(&User{}).CreateInBatches(list, BatchSize).Error
	return
}

// 根据主键查询数据
func (t *User) GetFirstByPk(db *gorm.DB, pk int) (err error, user User) {
	err = db.Model(t).First(&user, pk).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			err = ecode.DataNotExist
			return
		}
		return
	}

	return
}

func (t *User) GetUserByPhone(db *gorm.DB, phone string) (err error, user User) {
	err = db.Model(t).Where("phone = ?", phone).First(&user).Error

	if err != nil {
		return
	}

	return
}

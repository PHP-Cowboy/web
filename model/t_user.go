package model

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"web/utils/ecode"
)

// 用户
type User struct {
	Base
	Password string `gorm:"type:varchar(100);not null;comment:密码"`
	Name     string `gorm:"type:varchar(16);not null;comment:姓名"`
	RoleId   int    `gorm:"type:tinyint;not null;comment:角色表id"`
	Role     string `gorm:"type:varchar(16);not null;comment:角色(岗位)"`
	Status   int    `gorm:"type:tinyint;not null;default:1;comment:状态:0:未知,1:正常,2:禁用"`
}

type UserModel struct {
	DB *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		DB: db,
	}
}

// 保存用户
func (m *UserModel) Save(user *User) (err error) {
	err = m.DB.Model(&User{}).Save(user).Error
	return
}

// 批量创建用户
func (m *UserModel) CreateInBatches(list *[]User) (err error) {
	err = m.DB.Model(&User{}).CreateInBatches(list, BatchSize).Error
	return
}

// 批量创建或更新
func (m *UserModel) ReplaceCreateInBatches(list *[]User, values []string) (err error) {
	//[]string{"password", "name"}

	err = m.DB.Model(&User{}).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns(values),
		}).
		CreateInBatches(list, BatchSize).
		Error
	return
}

// 根据ID批量更新
func (m *UserModel) UpdateByIds(ids []int, mp map[string]interface{}) (err error) {
	err = m.DB.Model(&User{}).Where("id in (?)", ids).Updates(mp).Error

	return
}

// 根据主键查询数据
func (m *UserModel) GetFirstByPk(pk int) (err error, user User) {
	err = m.DB.Model(&User{}).First(&user, pk).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			err = ecode.DataNotExist
			return
		}
		return
	}

	return
}

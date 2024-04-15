package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 试题内容
type Question struct {
	Base
	CategoryId int    `gorm:"type:int(11);not null;comment:所属试题类别id"`
	Number     int    `gorm:"type:int(11);not null;comment:题目编号"`
	Topic      string `gorm:"type:varchar(255);not null;default:'';comment:题目"`
	A          string `gorm:"type:varchar(255);not null;default:'';comment:选项A"`
	B          string `gorm:"type:varchar(255);not null;default:'';comment:选项B"`
	C          string `gorm:"type:varchar(255);not null;default:'';comment:选项C"`
	D          string `gorm:"type:varchar(255);not null;default:'';comment:选项D"`
	E          string `gorm:"type:varchar(255);not null;default:'';comment:选项E"`
	Answer     string `gorm:"type:varchar(32);not null;default:'';comment:答案"`
	Analysis   string `gorm:"type:varchar(32);not null;default:'';comment:答案解析"`
}

func (t *Question) GetPageList(db *gorm.DB, form req.Id) (list []Question, err error) {

	err = db.Model(t).Where("category_id = ?", form.Id).Find(&list).Error
	return
}

func (t *Question) GetOneById(db *gorm.DB, id int) (content Question, err error) {
	err = db.Model(t).First(&content, id).Error
	return
}

func (t *Question) GetPrevById(db *gorm.DB, id int) (content Question, err error) {
	err = db.Model(t).Where("id < ? ", id).First(&content).Error
	return
}

func (t *Question) GetNextById(db *gorm.DB, id int) (content Question, err error) {
	err = db.Model(t).Where("id > ?", id).First(&content).Error
	return
}

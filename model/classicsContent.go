package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 典籍内容
type ClassicsContent struct {
	Base
	ClassicsId int    `gorm:"type:int(11);not null;comment:所属典籍id"`
	Title      string `gorm:"type:varchar(64);not null;comment:标题"`
	Content    string `gorm:"type:text;not null;comment:内容"`
}

func (t *ClassicsContent) GetOneById(db *gorm.DB, id int) (content ClassicsContent, err error) {
	err = db.Model(t).First(&content, id).Error
	return
}

func (t *ClassicsContent) GetPrevByIdClassicsId(db *gorm.DB, id, classicsId int) (content ClassicsContent, err error) {
	err = db.Model(t).Where("id < ? and classics_id = ?", id, classicsId).First(&content).Error
	return
}

func (t *ClassicsContent) GetNextByIdClassicsId(db *gorm.DB, id, classicsId int) (content ClassicsContent, err error) {
	err = db.Model(t).Where("id > ? and classics_id = ?", id, classicsId).First(&content).Error
	return
}

func (t *ClassicsContent) GetPageList(db *gorm.DB, form req.CatalogueList) (total int64, list []ClassicsContent, err error) {
	localDb := db.Model(t).Where("classics_id = ?", form.ClassicsId)

	err = localDb.Count(&total).Error

	if err != nil {
		return
	}

	err = localDb.Scopes(Paginate(form.Page, form.Size)).Find(&list).Error

	return
}

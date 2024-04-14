package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 伤寒论
type Typhoid struct {
	Base
	Title   string `gorm:"type:varchar(32);not null;comment:标题"`
	Brief   string `gorm:"type:varchar(64);not null;comment:简介"`
	Content string `gorm:"type:varchar(1024);not null;comment:内容"`
}

func (t *Typhoid) GetOneById(db *gorm.DB, id int) (content Typhoid, err error) {
	err = db.Model(t).First(&content, id).Error
	return
}

func (t *Typhoid) GetPrevById(db *gorm.DB, id int) (content Typhoid, err error) {
	err = db.Model(t).Where("id < ? ", id).First(&content).Error
	return
}

func (t *Typhoid) GetNextById(db *gorm.DB, id int) (content Typhoid, err error) {
	err = db.Model(t).Where("id > ?", id).First(&content).Error
	return
}

func (t *Typhoid) GetPageList(db *gorm.DB, form req.TyphoidCatalogueList) (total int64, list []Typhoid, err error) {
	localDb := db.Model(t).
		Where("title like ?", "%"+form.KeyWords+"%").
		Or("content like ?", "%"+form.KeyWords+"%")

	err = localDb.Count(&total).Error

	if err != nil {
		return
	}

	err = localDb.Scopes(Paginate(form.Page, form.Size)).Find(&list).Error

	return
}

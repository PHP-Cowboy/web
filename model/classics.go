package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 典籍
type Classics struct {
	Base
	CategoryId int    `gorm:"type:int(11);not null;comment:所属分类id"`
	Name       string `gorm:"type:varchar(32);not null;default:'';comment:名称"`
	Author     string `gorm:"type:varchar(32);not null;default:'';comment:作者"`
	Dynasty    string `gorm:"type:varchar(16);not null;default:'';comment:朝代"`
}

func (t *Classics) GetPageList(db *gorm.DB, form req.BookListByCategory) (total int64, list []Classics, err error) {
	localDb := db.Model(t).Where("category_id = ?", form.CategoryId)

	err = localDb.Count(&total).Error

	if err != nil {
		return
	}

	err = localDb.Scopes(Paginate(form.Page, form.Size)).Find(&list).Error

	return
}

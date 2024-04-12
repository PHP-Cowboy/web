package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 中医方剂
type Prescription struct {
	Base
	Title      string `gorm:"type:varchar(32);not null;comment:方剂标题"`
	Provenance string `gorm:"type:varchar(32);not null;default:'';comment:方剂出处"` //出处
	Content    string `gorm:"type:varchar(255);not null;default:'';comment:内容"`  //内容
	CategoryId int    `gorm:"type:int(11);not null;comment:所属中医方剂分类id"`          // 所属中医方剂分类id
}

func (t *Prescription) GetFirst(db *gorm.DB, form req.Id) (m Prescription, err error) {
	err = db.Model(t).First(&m, form.Id).Error
	return
}

func (t *Prescription) GetPageList(db *gorm.DB, form req.PrescriptionList) (total int64, list []Prescription, err error) {
	local := db.Model(t)

	local.Where(Prescription{CategoryId: form.CategoryId})

	if form.KeyWords != "" {
		local.Where("title like ?", "%"+form.KeyWords+"%").
			Or("provenance like ?", "%"+form.KeyWords+"%").
			Or("content like ?", "%"+form.KeyWords+"%")
	}

	err = local.Count(&total).Error

	if err != nil {
		return
	}

	err = db.Model(t).Scopes(Paginate(form.Page, form.Size)).Find(&list).Error
	return
}

// 中医方剂分类
type PrescriptionCategory struct {
	Base
	Name    string `gorm:"type:varchar(32);not null;comment:分类名称"`
	Picture string `gorm:"type:varchar(255);not null;default:'';comment:图片"` //图片
}

func (t *PrescriptionCategory) GetPageList(db *gorm.DB, form req.PrescriptionCategoryList) (total int64, list []PrescriptionCategory, err error) {
	local := db.Model(t)

	if form.KeyWords != "" {
		local.Where("name like ?", "%"+form.KeyWords+"%")
	}

	err = local.Count(&total).Error

	if err != nil {
		return
	}

	err = db.Model(t).Scopes(Paginate(form.Page, form.Size)).Find(&list).Error
	return
}

package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 常用方剂
type CommonlyPrescription struct {
	Base
	Title       string `gorm:"type:varchar(32);not null;comment:方剂标题"`
	Provenance  string `gorm:"type:varchar(32);not null;default:'';comment:方剂出处"`  //出处
	Constituent string `gorm:"type:varchar(255);not null;default:'';comment:组成成分"` //组成成分
	Usage       string `gorm:"type:varchar(255);not null;default:'';comment:用法"`   //用法
	Efficacy    string `gorm:"type:varchar(255);not null;default:'';comment:功效"`   //功效
	Mainly      string `gorm:"type:varchar(255);not null;default:'';comment:主治"`   //主治
	Song        string `gorm:"type:varchar(255);not null;default:'';comment:方歌"`   //方歌
	CategoryId  int    `gorm:"type:int(11);not null;comment:所属常用方剂分类id"`           // 所属常用方剂分类id
}

func (t *CommonlyPrescription) GetFirst(db *gorm.DB, form req.Id) (m CommonlyPrescription, err error) {
	err = db.Model(t).First(&m, form.Id).Error
	return
}

func (t *CommonlyPrescription) GetPageList(db *gorm.DB, form req.CommonlyPrescriptionList) (total int64, list []CommonlyPrescription, err error) {
	local := db.Model(t)

	local.Where(CommonlyPrescription{CategoryId: form.CategoryId})

	if form.KeyWords != "" {
		local.Where("title like ?", "%"+form.KeyWords+"%").
			Or("provenance like ?", "%"+form.KeyWords+"%").
			Or("constituent like ?", "%"+form.KeyWords+"%").
			Or("`usage` like ?", "%"+form.KeyWords+"%").
			Or("efficacy like ?", "%"+form.KeyWords+"%").
			Or("mainly like ?", "%"+form.KeyWords+"%").
			Or("song like ?", "%"+form.KeyWords+"%")
	}

	err = local.Count(&total).Error

	if err != nil {
		return
	}

	err = db.Model(t).Scopes(Paginate(form.Page, form.Size)).Find(&list).Error
	return
}

package model

import (
	"gorm.io/gorm"
)

// 方剂剂量表
type Dose struct {
	Base
	FormulaId int `gorm:"type:int(11);not null;comment:方剂id"`
	HerbId    int `gorm:"type:int(11);not null;comment:药物id"`
	Weight    int `gorm:"type:int(11);not null;comment:药量比例"`
}

func (t *Dose) CreateInBatches(db *gorm.DB, dataList []Dose) error {
	return db.Model(t).CreateInBatches(&dataList, BatchSize).Error
}

func (t *Dose) GetListByFormulaId(db *gorm.DB, formulaId int) (dataList []Dose, err error) {
	err = db.Model(t).Where("formula_id = ?", formulaId).Find(&dataList).Error

	return
}

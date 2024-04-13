package model

import (
	"gorm.io/gorm"
	"web/forms/req"
)

// 方剂图表
type PrescriptionGraph struct {
	Base
	PrescriptionId int    `gorm:"type:int(11);not null;comment:所属方剂id"`
	Type           int    `gorm:"type:tinyint(4);not null;comment:类型:1:症状群,2:舌象群,3:脉象群"`
	Name           string `gorm:"type:varchar(32);not null;default:'';comment:名称"`
	Num            int    `gorm:"type:int(11);not null;comment:数量"`
}

func (t *PrescriptionGraph) GetPageListByPrescriptionIdType(db *gorm.DB, form req.PrescriptionGraph) (list []PrescriptionGraph, err error) {
	err = db.Model(t).Where(PrescriptionGraph{PrescriptionId: form.PrescriptionId, Type: form.Type}).Find(&list).Error

	return
}

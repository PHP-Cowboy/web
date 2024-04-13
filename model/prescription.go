package model

// 方剂
type Prescription struct {
	Base
}

// 名家
type Celebrity struct {
	Base
	NotableDoctor  string `gorm:"type:varchar(16);not null;default:'';comment:医家"`
	Provenance     string `gorm:"type:varchar(32);not null;default:'';comment:出处"`
	Content        string `gorm:"type:varchar(255);not null;default:'';comment:内容"`
	PrescriptionId int    `gorm:"type:int(11);not null;comment:所属方剂id"`
}

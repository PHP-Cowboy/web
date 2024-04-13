package model

// 药物
type Drug struct {
	Base
	Name         string `gorm:"type:varchar(32);not null;comment:名称"`
	Abbreviation string `gorm:"type:varchar(32);not null;comment:名称缩写"`
	Describe     string `gorm:"type:varchar(255);not null;comment:性味归经"`
	Efficacy     string `gorm:"type:varchar(255);not null;comment:功效,主治"`
}

// 用户自定义药物
type UserDrug struct {
	Base
	Name         string `gorm:"type:varchar(32);not null;comment:名称"`
	Abbreviation string `gorm:"type:varchar(32);not null;comment:名称缩写"`
	Describe     string `gorm:"type:varchar(255);not null;comment:性味归经"`
	Efficacy     string `gorm:"type:varchar(255);not null;comment:功效,主治"`
}

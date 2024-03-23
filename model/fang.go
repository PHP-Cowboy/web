package model

type Fang struct {
	Id        int    `gorm:"primaryKey;type:int(11) unsigned AUTO_INCREMENT;comment:id" json:"id"`
	Mingcheng string `gorm:"type:varchar(255);comment:fang" json:"mingcheng"`
	Jiexi     string `gorm:"type:varchar(255);comment:jiexi" json:"jiexi"`
}

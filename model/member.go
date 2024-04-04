package model

// 会员
type Member struct {
	Base
	Uid    int `gorm:"type:int(11);not null;comment:用户id"`
	Level  int `gorm:"type:tinyint(4);not null;comment:级别:1:永久,2:定期"`
	Expire int `gorm:"type:int(11);not null;default:0;comment:过期时间"` //过期时间
}

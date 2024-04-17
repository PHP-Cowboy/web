package model

import (
	"gorm.io/gorm"
	"time"
)

// 订单
type Order struct {
	Base
	Uid          int     `gorm:"type:int(11);not null;comment:用户id"`
	OrderNo      string  `gorm:"type:varchar(64);not null;unique;default:'';comment:订单编号"`
	Status       int     `gorm:"type:tinyint(4);not null;comment:状态:1:未支付,2:支付完成"`
	GiftId       int     `gorm:"type:int(11);default:0;comment:购买礼包id"`
	Price        float64 `gorm:"type:decimal(10,2);comment:价格"`
	PayType      int     `gorm:"type:tinyint(4);not null;comment:支付方式:1:支付宝,2:微信"`
	CompleteTime int     `gorm:"type:int(11);default:0;comment:完成时间"`
}

const (
	OrderStatus       = iota
	OrderStatusNotPay //未支付
	OrderStatusPaid   //已支付
)

func (t *Order) Save(db *gorm.DB) (err error) {
	err = db.Model(t).Create(t).Error
	return
}

func (t *Order) OrderPay(db *gorm.DB, orderNo string) (err error) {
	err = db.Model(t).
		Where("order_no = ?", orderNo).
		Updates(map[string]interface{}{
			"status":        OrderStatusPaid,
			"complete_time": time.Now().Unix(),
		}).Error

	return
}

func (t *Order) GetFirstById(db *gorm.DB, id int) (data Order, err error) {
	err = db.Model(t).Where("id = ?", id).First(&data).Error
	return
}

func (t *Order) GetList(db *gorm.DB) (list []Order, err error) {
	err = db.Model(t).Where(t).Find(&list).Error
	return
}

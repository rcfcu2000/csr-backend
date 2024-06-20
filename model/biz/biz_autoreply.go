package models

import "time"

type BizAutoReply struct {
	ID         uint      `gorm:"primaryKey"`
	Question   string    `gorm:"size:355"`
	Answer     string    `gorm:"size:555"`
	RegEx      string    `gorm:"size:555"`
	UpdateTime time.Time `gorm:"default:NULL"`
	UpdatedBy  string    `gorm:"size:255"`
	Status     int       `gorm:"not null; default 1"`
	ShopId     uint      `json:"shopId" gorm:"NOT NULL;int"`
}

func (BizAutoReply) TableName() string {
	return "biz_auto_reply"
}

package models

import "time"

type BizQaQuestion struct {
	Qid        uint      `gorm:"primaryKey;autoIncrement:false"`
	Question   string    `gorm:"size:355"`
	UpdateTime time.Time `gorm:"default:NULL"`
	UpdatedBy  string    `gorm:"size:255"`
}

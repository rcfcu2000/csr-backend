package models

import "time"

type BizQuestionType struct {
	Qid        uint      `gorm:"primaryKey;autoIncrement:false"`
	TypeID     uint      `gorm:"primaryKey;autoIncrement:false"`
	UpdateTime time.Time `gorm:"default:NULL"`
	UpdatedBy  string    `gorm:"size:255"`
}

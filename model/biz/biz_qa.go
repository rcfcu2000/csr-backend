package models

import "time"

type BizQa struct {
	ID         uint      `gorm:"primaryKey"`
	Question   string    `gorm:"size:355"`
	Answer     string    `gorm:"size:555"`
	UpdateTime time.Time `gorm:"default:NULL"`
	UpdatedBy  string    `gorm:"size:255"`
}

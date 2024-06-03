package models

import "time"

type BizMerchant struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:255"`
	Alias       string    `gorm:"size:255"`
	PictureLink string    `gorm:"size:255"`
	UpdateTime  time.Time `gorm:"default:NULL"`
	UpdatedBy   string    `gorm:"size:255"`
}

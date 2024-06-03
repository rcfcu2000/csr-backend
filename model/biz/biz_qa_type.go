package models

type BizQaType struct {
	ID    uint   `gorm:"primaryKey"`
	QType string `gorm:"size:255"`
}

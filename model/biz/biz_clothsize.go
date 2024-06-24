package models

import "time"

type BizClothSize struct {
	ID         uint        `gorm:"not null;unique;primary_key;autoIncrement:true;"`
	Name       string      `json:"name" gorm:"size:255;unique"`
	CategoryID uint        `json:"categoryId" gorm:"not null"`
	Category   BizCategory `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	UpdateTime time.Time   `gorm:"default:NULL"`
	UpdatedBy  string      `gorm:"size:255"`
	RefCount   int         `gorm:"default 0;not null"`
	Status     int         `gorm:"not null; default 1"`
	SizeInfo   string      `gorm:"type:varchar(5000);"`
	ShopId     uint        `json:"shopId" gorm:"NOT NULL;int"`
}

func (BizClothSize) TableName() string {
	return "biz_cloth_size_info"
}

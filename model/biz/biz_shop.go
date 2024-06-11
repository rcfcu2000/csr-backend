package models

type BizShop struct {
	ID              uint        `gorm:"not null;unique;primary_key;autoIncrement:true;"`
	Name            string      `gorm:"size:255"`
	Nickname        string      `gorm:"size:255"`
	CategoryID      uint        `gorm:"not null"`
	Category        BizCategory `gorm:"foreignKey:CategoryID;references:ID"`
	BrandManagement string      `gorm:"size:355"`
	BrandBelief     string      `gorm:"size:355"`
	BrandAdvantage  string      `gorm:"size:355"`
	BrandInfo       string      `gorm:"size:355"`
}

func (BizShop) TableName() string {
	return "biz_shop"
}

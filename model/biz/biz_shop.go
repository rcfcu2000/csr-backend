package models

type BizShop struct {
	ID              uint        `gorm:"not null;unique;primary_key;autoIncrement:true;"`
	Name            string      `json:"name" gorm:"size:255;unique"`
	Nickname        string      `json:"nickName" gorm:"size:255"`
	CategoryID      uint        `json:"categoryId" gorm:"not null"`
	Category        BizCategory `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	BrandManagement string      `json:"brandManagement" gorm:"size:355"`
	BrandBelief     string      `json:"brandBelief" gorm:"size:355"`
	BrandAdvantage  string      `json:"brandAdvantage" gorm:"size:355"`
	RegEx           string      `json:"regEx" gorm:"size:555"`
	BrandInfo       string      `json:"brandInfo" gorm:"size:555"`
}

func (BizShop) TableName() string {
	return "biz_shop"
}

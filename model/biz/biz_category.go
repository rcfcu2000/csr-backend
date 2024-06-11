package models

type BizCategory struct {
	ID       uint   `gorm:"not null;unique;primary_key;autoIncrement:true;"`
	Name     string `gorm:"size:255"`
	Industry string `gorm:"size:255"`
}

func (BizCategory) TableName() string {
	return "biz_category"
}

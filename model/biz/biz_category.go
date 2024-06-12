package models

type BizCategory struct {
	ID         uint   `gorm:"not null;unique;primary_key;autoIncrement:true;"`
	Name       string `gorm:"size:255;unique"`
	Industry   string `gorm:"size:255"`
	ParentID   uint   `json:"parentId" gorm:"not null"`
	Level      uint   `gorm:"not null; default 1"`
	Status     uint   `gorm:"not null; default 1"`
	PresetText string `json:"presetText" gorm:"defaul null"`
}

func (BizCategory) TableName() string {
	return "biz_category"
}

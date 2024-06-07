package models

import "time"

type BizLinks struct {
	CreatedAt time.Time  // 创建时间
	DeletedAt *time.Time `sql:"index"`
	LinkId    uint       `json:"linkId" gorm:"not null;unique;primary_key;autoIncrement:true;comment:链接ID"`
	Link      string     `json:"name" gorm:"size:255;comment:商品链接"`
	TaobaoId  uint       `json:"taobaoId" gorm:"comment:淘宝ID"`
	UpdatedAt time.Time  // 更新时间
	UpdatedBy string     `json:"updatedBy" gorm:"size:255;comment:修改人"`
}

func (BizLinks) TableName() string {
	return "biz_links"
}

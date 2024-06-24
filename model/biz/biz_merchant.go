package models

import "time"

type BizMerchant struct {
	CreatedAt     time.Time  // 创建时间
	DeletedAt     *time.Time `sql:"index"`
	MerchantID    uint       `json:"merchantId" gorm:"not null;unique;primary_key;autoIncrement:true;comment:商品ID"` // 商品ID
	Name          string     `json:"name" gorm:"size:255;comment:标题"`
	Alias         string     `json:"alias" gorm:"size:255;comment:简称"`
	Information   string     `json:"info" gorm:"size:555;comment:商品信息"`
	PictureLink   string     `json:"pictureLink" gorm:"size:255;comment:封面图"`
	UpdatedAt     time.Time  // 更新时间
	UpdatedBy     string     `json:"updatedBy" gorm:"size:255;comment:修改人"`
	MerchantLinks []BizLinks `json:"merchantLinks" gorm:"many2many:biz_merchant_links"`
	SizeInfoId    int        `json:"sizeinfoId" gorm:"NOT NULL;int"`
	ShopId        int        `json:"shopId" gorm:"NOT NULL;int"`
}

func (BizMerchant) TableName() string {
	return "biz_merchants"
}

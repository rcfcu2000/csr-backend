package models

// BizMerchantLinks 是 bizMerchant 和 bizLinks 的连接表
type BizMerchantLinks struct {
	BizMerchantId uint `gorm:"column:biz_merchant_merchant_id"`
	BizLinkId     uint `gorm:"column:biz_links_link_id"`
}

func (s *BizMerchantLinks) TableName() string {
	return "biz_merchant_links"
}

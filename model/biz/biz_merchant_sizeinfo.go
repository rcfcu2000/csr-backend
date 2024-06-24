package models

// BizMerchantSizeInfo 是 bizMerchant 和 bizClothSize 的连接表
type BizMerchantSizeInfo struct {
	BizMerchantId uint `gorm:"column:biz_merchant_merchant_id"`
	BizSizeInfoId uint `gorm:"column:biz_clothsize_sizeinfo_id"`
}

type UpdateMList struct {
	ClothSizeInfoId int   `json:"cloth_size_info_id"`
	MerchantIds     []int `json:"merchant_ids"`
}

func (s *BizMerchantSizeInfo) TableName() string {
	return "biz_merchant_sizeinfo"
}

package biz

import (
	"errors"
	"xtt/global"
	"xtt/model/biz"
	"xtt/model/common/request"

	"gorm.io/gorm"
)

type MerchantService struct{}

// CreateMerchants creates multiple merchants in the database
func (exa *MerchantService) CreateMerchants(merchants []models.BizMerchant) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, merchant := range merchants {
			if err := tx.Create(&merchant).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// CreateMerchant creates a new merchant in the database
func (exa *MerchantService) CreateMerchant(merchant *models.BizMerchant) error {
	if err := global.GVA_DB.Create(merchant).Error; err != nil {
		return err
	}
	return nil
}

// GetMerchant retrieves a merchant by ID from the database
func (exa *MerchantService) GetMerchant(id string) (*models.BizMerchant, error) {
	var merchant models.BizMerchant
	if err := global.GVA_DB.Preload("MerchantLinks").First(&merchant, "merchant_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("merchant not found")
		}
		return nil, err
	}
	return &merchant, nil
}

// GetMerchant retrieves merchants by taobaoid from the database
func (exa *MerchantService) GetMerchantsByTid(id string) ([]*models.BizMerchant, error) {
	var merchants []*models.BizMerchant

	sql_string := "select * from biz_merchants where merchant_id in (select biz_merchant_merchant_id from biz_merchant_links where biz_links_link_id in (select link_id from biz_links where taobao_id = ?))"
	err := global.GVA_DB.Raw(sql_string, id).Scan(&merchants).Error

	if err != nil {
		return nil, err
	}

	return merchants, nil
}

// @function: GetMerchantList
// @description: 分页获取数据
// @param: info request.PageInfo
// @return: err error, list interface{}, total int64
func (userService *MerchantService) GetMerchantList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	shopId := info.ShopId
	db := global.GVA_DB.Model(&models.BizMerchant{}).Where("shop_id = ?", shopId)
	var merchantList []models.BizMerchant
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("MerchantLinks").Find(&merchantList).Error
	return merchantList, total, err
}

// UpdateMerchant updates an existing merchant in the database
func (exa *MerchantService) UpdateMerchant(id string, updatedMerchant *models.BizMerchant) error {
	var merchant models.BizMerchant
	if err := global.GVA_DB.First(&merchant, "merchant_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("merchant not found")
		}
		return err
	}

	if err := global.GVA_DB.Model(&merchant).Updates(updatedMerchant).Error; err != nil {
		return err
	}
	return nil
}

// DeleteMerchant deletes a merchant by ID from the database
func (exa *MerchantService) DeleteMerchant(id string) error {
	if err := global.GVA_DB.Delete(&models.BizMerchant{}, "merchant_id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

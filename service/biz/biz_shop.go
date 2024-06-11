package biz

import (
	"fmt"
	"xtt/global"
	"xtt/model/biz"
)

type BizShopService struct {
}

func (s *BizShopService) CreateShop(shop *models.BizShop) error {
	shop.BrandInfo = s.calculateBrandInfo(shop)
	return global.GVA_DB.Create(shop).Error
}

func (s *BizShopService) GetShopByID(id uint) (*models.BizShop, error) {
	var shop models.BizShop
	if err := global.GVA_DB.Preload("Category").First(&shop, id).Error; err != nil {
		return nil, err
	}
	return &shop, nil
}

func (s *BizShopService) UpdateShop(shop *models.BizShop) error {
	shop.BrandInfo = s.calculateBrandInfo(shop)
	return global.GVA_DB.Save(shop).Error
}

func (s *BizShopService) DeleteShop(id uint) error {
	return global.GVA_DB.Delete(&models.BizShop{}, id).Error
}

func (s *BizShopService) ListShops() ([]models.BizShop, error) {
	var shops []models.BizShop
	if err := global.GVA_DB.Preload("Category").Find(&shops).Error; err != nil {
		return nil, err
	}
	return shops, nil
}

func (s *BizShopService) calculateBrandInfo(shop *models.BizShop) string {
	return fmt.Sprintf("品牌理念: %s; 品牌优势: %s", shop.BrandAdvantage, shop.BrandBelief)
}

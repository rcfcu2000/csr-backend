package biz

import (
	"xtt/global"
	"xtt/model/biz"
)

type BizShopService struct {
}

func (s *BizShopService) CreateShop(shop *models.BizShop) error {
	return global.GVA_DB.Create(shop).Error
}

func (s *BizShopService) GetShopByID(id uint) (*models.BizShop, error) {
	var shop models.BizShop
	if err := global.GVA_DB.Preload("Category").First(&shop, id).Error; err != nil {
		return nil, err
	}
	return &shop, nil
}

func (s *BizShopService) GetShopByName(name string) (*models.BizShop, error) {
	var shop models.BizShop
	if err := global.GVA_DB.Preload("Category").Where("name = ?", name).First(&shop).Error; err != nil {
		return nil, err
	}
	return &shop, nil
}

func (s *BizShopService) UpdateShop(shop *models.BizShop) error {
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

func (s *BizShopService) GetOrCreateCategory(category *models.BizCategory) error {
	if category.ID <= 0 {
		if err := global.GVA_DB.Create(category).Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *BizShopService) ListCategories() ([]models.BizCategory, error) {
	var categories []models.BizCategory
	if err := global.GVA_DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

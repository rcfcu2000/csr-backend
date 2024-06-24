package biz

import (
	"time"

	"xtt/global"
	models "xtt/model/biz"
	"xtt/model/common/request"

	"gorm.io/gorm"
)

type BizClothSizeService struct{}

func (s *BizClothSizeService) CreateClothSizeInfo(bizQa *models.BizClothSize) error {
	return global.GVA_DB.Create(bizQa).Error
}

func (s *BizClothSizeService) GetClothSizeInfoByID(id uint) (*models.BizClothSize, error) {
	var bizQa models.BizClothSize
	if err := global.GVA_DB.First(&bizQa, id).Error; err != nil {
		return nil, err
	}
	return &bizQa, nil
}

func (s *BizClothSizeService) GetClothSizeInfoByMerchat(merchant_id int, shopid int) (*models.BizClothSize, error) {
	var bizQa models.BizClothSize

	if merchant_id <= 0 {

		if err := global.GVA_DB.Where("status = 2 and shop_id = ?", shopid).Find(&bizQa).Error; err != nil {
			return nil, err
		}
		return &bizQa, nil
	} else {

		if err := global.GVA_DB.Model(&models.BizClothSize{}).Where("id = (select biz_clothsize_sizeinfo_id from biz_merchant_sizeinfo where  biz_merchant_merchant_id = ?) and shop_id = ?", merchant_id, shopid).First(&bizQa).Error; err != nil {
			return nil, err
		}
		return &bizQa, nil
	}
}

func (s *BizClothSizeService) UpdateClothSizeInfo(bizQa *models.BizClothSize) error {
	bizQa.UpdateTime = time.Now()
	err := global.GVA_DB.Save(bizQa).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *BizClothSizeService) GetMerchantList(id uint) ([]models.BizMerchantSizeInfo, error) {
	var mlist []models.BizMerchantSizeInfo
	db := global.GVA_DB.Model(&models.BizMerchantSizeInfo{}).Where("biz_clothsize_sizeinfo_id = ?", id)
	err := db.Find(&mlist).Error
	return mlist, err
}

func (s *BizClothSizeService) UpdateMerchantList(mlist *models.UpdateMList) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {

		if err := global.GVA_DB.Model(&models.BizMerchantSizeInfo{}).Where("biz_clothsize_sizeinfo_id = ?", mlist.ClothSizeInfoId).Delete(&models.BizMerchantSizeInfo{}).Error; err != nil {
			return err
		}

		for _, merchantid := range mlist.MerchantIds {
			var ms models.BizMerchantSizeInfo
			ms.BizSizeInfoId = uint(mlist.ClothSizeInfoId)
			ms.BizMerchantId = uint(merchantid)
			if err := tx.Create(&ms).Error; err != nil {
				return err
			}
		}

		refCount := len(mlist.MerchantIds)
		if err := global.GVA_DB.Model(&models.BizClothSize{}).Where("id = ?", mlist.ClothSizeInfoId).Update("ref_count", refCount).Error; err != nil {
			return err
		}

		return nil
	})
}

func (s *BizClothSizeService) DeleteClothSizeInfo(id uint) error {

	// // Decrement the RefCount for each related BizQaType
	// for _, bqt := range bizQuestionTypes {
	// 	if err := global.GVA_DB.Model(&models.BizQaType{}).Where("id = ?", bqt.TypeID).Update("ref_count", gorm.Expr("ref_count - ?", 1)).Error; err != nil {
	// 		return err
	// 	}
	// }

	return global.GVA_DB.Delete(&models.BizClothSize{}, id).Error
}

func (s *BizClothSizeService) GetClothSizeInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	shopId := info.ShopId
	db := global.GVA_DB.Model(&models.BizClothSize{}).Preload("Category").Where("shop_id = ?", shopId)
	var arList []models.BizClothSize
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&arList).Error
	return arList, total, err
}

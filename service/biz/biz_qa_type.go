package biz

import (
	"xtt/global"
	"xtt/model/biz"
	"xtt/model/common/request"

	"gorm.io/gorm"
)

type BizQaTypeService struct{}

func (s *BizQaTypeService) CreateBizQaType(bizQaType *models.BizQaType) error {
	return global.GVA_DB.Create(bizQaType).Error
}

func (s *BizQaTypeService) GetBizQaTypeByID(id uint) (*models.BizQaType, error) {
	var bizQaType models.BizQaType
	if err := global.GVA_DB.First(&bizQaType, id).Error; err != nil {
		return nil, err
	}
	return &bizQaType, nil
}

func (s *BizQaTypeService) UpdateBizQaType(bizQaType *models.BizQaType) error {
	return global.GVA_DB.Save(bizQaType).Error
}

func (s *BizQaTypeService) DeleteBizQaType(id uint) error {
	return global.GVA_DB.Delete(&models.BizQaType{}, id).Error
}

func (s *BizQaTypeService) GetQaTypeList(info request.PageInfo, kb_type int) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&models.BizQaType{}).Where("kb_type = ?", kb_type)
	var qaTypeList []models.BizQaType
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&qaTypeList).Error
	return qaTypeList, total, err
}

func (s *BizQaTypeService) IncrementRefCount(id uint) error {
	return global.GVA_DB.Model(&models.BizQaType{}).Where("id = ?", id).Update("ref_count", gorm.Expr("ref_count + ?", 1)).Error
}

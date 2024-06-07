package biz

import (
	"time"

	"xtt/global"
	"xtt/model/biz"
	"xtt/model/common/request"
)

type BizQaService struct{}

func (s *BizQaService) CreateBizQa(bizQa *models.BizQa) error {
	return global.GVA_DB.Create(bizQa).Error
}

func (s *BizQaService) GetBizQaByID(id uint) (*models.BizQa, error) {
	var bizQa models.BizQa
	if err := global.GVA_DB.First(&bizQa, id).Error; err != nil {
		return nil, err
	}
	return &bizQa, nil
}

func (s *BizQaService) UpdateBizQa(bizQa *models.BizQa) error {
	bizQa.UpdateTime = time.Now()
	return global.GVA_DB.Save(bizQa).Error
}

func (s *BizQaService) DeleteBizQa(id uint) error {
	return global.GVA_DB.Delete(&models.BizQa{}, id).Error
}

func (s *BizQaService) GetQaList(info request.PageInfo, kb_type int) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&models.BizQa{}).Where("kb_type = ?", kb_type)
	var qaList []models.BizQa
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&qaList).Error
	return qaList, total, err
}

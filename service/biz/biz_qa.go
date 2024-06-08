package biz

import (
	"time"

	"xtt/global"
	"xtt/model/biz"
	"xtt/model/common/request"

	"gorm.io/gorm"
)

type BizQaService struct{}

func (s *BizQaService) CreateBizQa(bizQa *models.BizQa) error {
	return global.GVA_DB.Create(bizQa).Error
}

func (s *BizQaService) GetBizQaByID(id uint) (*models.BizQa, error) {
	var bizQa models.BizQa
	if err := global.GVA_DB.Preload("QaTypes").First(&bizQa, id).Error; err != nil {
		return nil, err
	}
	return &bizQa, nil
}

func (s *BizQaService) GetBizQaByQuestion(question string) (*models.BizQa, error) {
	var bizQa models.BizQa
	if err := global.GVA_DB.Preload("QaTypes").Where("question = ?", question).First(&bizQa).Error; err != nil {
		return nil, err
	}
	return &bizQa, nil
}

func (s *BizQaService) UpdateBizQa(bizQa *models.BizQa) error {
	bizQa.UpdateTime = time.Now()
	err := global.GVA_DB.Save(bizQa).Error
	if err != nil {
		return err
	}

	if bizQa.QaTypes != nil {
		// Find related biz_question_types entries
		var bizQuestionTypes []models.BizQuestionType
		if err := global.GVA_DB.Where("biz_qa_id = ?", bizQa.ID).Find(&bizQuestionTypes).Error; err != nil {
			return err
		}

		// Decrement the RefCount for each related BizQaType
		for _, bqt := range bizQuestionTypes {
			if err := global.GVA_DB.Model(&models.BizQaType{}).Where("id = ?", bqt.TypeID).Update("ref_count", gorm.Expr("ref_count - ?", 1)).Error; err != nil {
				return err
			}
		}

		for _, qatype := range bizQa.QaTypes {
			if err := global.GVA_DB.Model(&models.BizQaType{}).Where("id = ?", qatype.ID).Update("ref_count", gorm.Expr("ref_count + ?", 1)).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *BizQaService) DeleteBizQa(id uint) error {
	// Find related biz_question_types entries
	var bizQuestionTypes []models.BizQuestionType
	if err := global.GVA_DB.Where("biz_qa_id = ?", id).Find(&bizQuestionTypes).Error; err != nil {
		return err
	}

	// Delete the biz_qa record
	if err := global.GVA_DB.Delete(&models.BizQa{}, id).Error; err != nil {
		return err
	}

	// Decrement the RefCount for each related BizQaType
	for _, bqt := range bizQuestionTypes {
		if err := global.GVA_DB.Model(&models.BizQaType{}).Where("id = ?", bqt.TypeID).Update("ref_count", gorm.Expr("ref_count - ?", 1)).Error; err != nil {
			return err
		}
	}

	if err := global.GVA_DB.Delete(&models.BizQuestionType{}, "biz_qa_id = ?", id).Error; err != nil {
		return err
	}
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
	err = db.Limit(limit).Offset(offset).Preload("QaTypes").Find(&qaList).Error
	return qaList, total, err
}

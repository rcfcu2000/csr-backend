package biz

import (
	"time"

	"xtt/global"
	models "xtt/model/biz"
	"xtt/model/common/request"
)

type BizAutoReplyService struct{}

func (s *BizAutoReplyService) CreateAutoReply(bizQa *models.BizAutoReply) error {
	return global.GVA_DB.Create(bizQa).Error
}

func (s *BizAutoReplyService) GetBizQaByID(id uint) (*models.BizAutoReply, error) {
	var bizQa models.BizAutoReply
	if err := global.GVA_DB.First(&bizQa, id).Error; err != nil {
		return nil, err
	}
	return &bizQa, nil
}

func (s *BizAutoReplyService) GetBizQaByQuestion(questions []string, shopid int) ([]models.BizQa, error) {
	var bizQa []models.BizQa
	if err := global.GVA_DB.Preload("QaTypes").Where("question in (?) and shop_id = ?", questions, shopid).Find(&bizQa).Error; err != nil {
		return nil, err
	}
	return bizQa, nil
}

func (s *BizAutoReplyService) UpdateAutoReply(bizQa *models.BizAutoReply) error {
	bizQa.UpdateTime = time.Now()
	err := global.GVA_DB.Save(bizQa).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *BizAutoReplyService) DeleteAutoReply(id uint) error {

	// // Decrement the RefCount for each related BizQaType
	// for _, bqt := range bizQuestionTypes {
	// 	if err := global.GVA_DB.Model(&models.BizQaType{}).Where("id = ?", bqt.TypeID).Update("ref_count", gorm.Expr("ref_count - ?", 1)).Error; err != nil {
	// 		return err
	// 	}
	// }

	return global.GVA_DB.Delete(&models.BizAutoReply{}, id).Error
}

func (s *BizAutoReplyService) GetAutoReplyList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	shopId := info.ShopId
	db := global.GVA_DB.Model(&models.BizAutoReply{}).Where("shop_id = ?", shopId)
	var arList []models.BizAutoReply
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&arList).Error
	return arList, total, err
}

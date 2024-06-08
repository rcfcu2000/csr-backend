package models

type BizQuestionType struct {
	Qid    uint `gorm:"column:biz_qa_id"`
	TypeID uint `gorm:"column:biz_qa_type_id"`
}

func (BizQuestionType) TableName() string {
	return "biz_question_types"
}

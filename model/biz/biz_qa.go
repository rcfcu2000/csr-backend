package models

import "time"

type BizQa struct {
	ID         uint        `gorm:"primaryKey"`
	Question   string      `gorm:"size:355"`
	Answer     string      `gorm:"size:555"`
	UpdateTime time.Time   `gorm:"default:NULL"`
	UpdatedBy  string      `gorm:"size:255"`
	KbType     int         `gorm:"NOT NULL;int"` // 1 通用知识库   2 定制知识库
	QaTypes    []BizQaType `json:"qa_types" gorm:"many2many:biz_question_types"`
}

func (BizQa) TableName() string {
	return "biz_qas"
}

package models

type BizQaType struct {
	ID       uint   `gorm:"primaryKey"`
	QType    string `gorm:"size:255"`
	KbType   int    `gorm:"NOT NULL;int"` // 1 通用知识库   2 定制知识库
	RefCount int    `gorm:"NOT NULL;int"`
}

func (BizQaType) TableName() string {
	return "biz_qa_types"
}

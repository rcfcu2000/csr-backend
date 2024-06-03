package models

import "time"

type BizMessage struct {
	MTime      time.Time `json:"m_time" gorm:"primaryKey"`
	Direction  int64     `json:"direction" gorm:"primaryKey"`
	UserNick   string    `json:"user_nick" gorm:"primaryKey"`
	CsrNick    string    `json:"csr_nick"`
	Content    string    `json:"content,omitempty"`
	UrlLink    string    `json:"url_link,omitempty"`
	TemplateID int64     `json:"template_id,omitempty"`
}

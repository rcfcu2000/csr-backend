package models

import "time"

type BizMessage struct {
	MTime      time.Time `gorm:"primaryKey"`
	Direction  int       `gorm:"primaryKey"`
	UserNick   string    `gorm:"size:255"`
	CsrNick    string    `gorm:"size:255"`
	Content    string    `gorm:"size:555"`
	UrlLink    string    `gorm:"size:255"`
	TemplateID int
}

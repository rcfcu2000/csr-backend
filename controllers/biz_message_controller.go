package controllers

import (
	"csr-backend/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BizMessageController struct {
	DB *gorm.DB
}

func NewBizMessageController(db *gorm.DB) *BizMessageController {
	return &BizMessageController{DB: db}
}

type BizMessage struct {
	MTime      string `json:"m_time" binding:"required"`
	Direction  int64  `json:"direction" binding:"required"`
	UserNick   string `json:"user_nick" binding:"required"`
	CsrNick    string `json:"csr_nick" binding:"required"`
	Content    string `json:"content"`
	UrlLink    string `json:"url_link"`
	TemplateID int64  `json:"template_id"`
}

type CreateBizMessagesRequest struct {
	Messages []BizMessage `json:"messages" binding:"required"`
}

// CreateBizMessages godoc
// @Summary Create multiple new BizMessages
// @Description Create multiple new BizMessage entries
// @Tags biz_message
// @Accept json
// @Produce json
// @Param data body CreateBizMessagesRequest true "BizMessages"
// @Success 200 {array} models.BizMessage
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/biz_messages [post]
func (ctrl *BizMessageController) CreateBizMessages(c *gin.Context) {
	var req CreateBizMessagesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	var bizMessages []models.BizMessage

	for _, message := range req.Messages {
		// Parse the date-time string into a time.Time object
		layout := "2006-01-02 15:04:05"
		mTime, err := time.Parse(layout, message.MTime)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid date format"})
			return
		}

		bizMessage := models.BizMessage{
			MTime:      mTime,
			Direction:  message.Direction,
			UserNick:   message.UserNick,
			CsrNick:    message.CsrNick,
			Content:    message.Content,
			UrlLink:    message.UrlLink,
			TemplateID: message.TemplateID,
		}

		bizMessages = append(bizMessages, bizMessage)
	}

	if err := ctrl.DB.Save(&bizMessages).Error; err != nil {
		fmt.Println("Failed to save records")
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to save records"})
		return
	}

	c.JSON(http.StatusOK, bizMessages)
}

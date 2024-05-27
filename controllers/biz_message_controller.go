package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"your_project/models"
)

type BizMessageController struct {
	DB *gorm.DB
}

func NewBizMessageController(db *gorm.DB) *BizMessageController {
	return &BizMessageController{DB: db}
}

func (ctrl *BizMessageController) CreateBizMessage(c *gin.Context) {
	var newMessage models.BizMessage
	if err := c.ShouldBindJSON(&newMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.DB.Create(&newMessage)
	c.JSON(http.StatusOK, newMessage)
}

func (ctrl *BizMessageController) GetBizMessage(c *gin.Context) {
	mTime, err := time.Parse(time.RFC3339, c.Param("m_time"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format"})
		return
	}
	direction := c.Param("direction")
	var message models.BizMessage
	if err := ctrl.DB.First(&message, "m_time = ? AND direction = ?", mTime, direction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, message)
}

func (ctrl *BizMessageController) UpdateBizMessage(c *gin.Context) {
	mTime, err := time.Parse(time.RFC3339, c.Param("m_time"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format"})
		return
	}
	direction := c.Param("direction")
	var message models.BizMessage
	if err := ctrl.DB.First(&message, "m_time = ? AND direction = ?", mTime, direction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.DB.Save(&message)
	c.JSON(http.StatusOK, message)
}

func (ctrl *BizMessageController) DeleteBizMessage(c *gin.Context) {
	mTime, err := time.Parse(time.RFC3339, c.Param("m_time"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format"})
		return
	}
	direction := c.Param("direction")
	if err := ctrl.DB.Delete(&models.BizMessage{}, "m_time = ? AND direction = ?", mTime, direction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted"})
}

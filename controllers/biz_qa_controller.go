package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"your_project/models"
)

type BizQaController struct {
	DB *gorm.DB
}

func NewBizQaController(db *gorm.DB) *BizQaController {
	return &BizQaController{DB: db}
}

func (ctrl *BizQaController) CreateBizQa(c *gin.Context) {
	var newQa models.BizQa
	if err := c.ShouldBindJSON(&newQa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newQa.UpdateTime = time.Now()
	ctrl.DB.Create(&newQa)
	c.JSON(http.StatusOK, newQa)
}

func (ctrl *BizQaController) GetBizQa(c *gin.Context) {
	id := c.Param("id")
	var qa models.BizQa
	if err := ctrl.DB.First(&qa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, qa)
}

func (ctrl *BizQaController) UpdateBizQa(c *gin.Context) {
	id := c.Param("id")
	var qa models.BizQa
	if err := ctrl.DB.First(&qa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	if err := c.ShouldBindJSON(&qa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	qa.UpdateTime = time.Now()
	ctrl.DB.Save(&qa)
	c.JSON(http.StatusOK, qa)
}

func (ctrl *BizQaController) DeleteBizQa(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.DB.Delete(&models.BizQa{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted"})
}

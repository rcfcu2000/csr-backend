package controllers

import (
	"net/http"

	"csr-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BizQaTypeController struct {
	DB *gorm.DB
}

func NewBizQaTypeController(db *gorm.DB) *BizQaTypeController {
	return &BizQaTypeController{DB: db}
}

func (ctrl *BizQaTypeController) CreateBizQaType(c *gin.Context) {
	var newQaType models.BizQaType
	if err := c.ShouldBindJSON(&newQaType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.DB.Create(&newQaType)
	c.JSON(http.StatusOK, newQaType)
}

// GetBizQaType godoc
// @Summary Get a BizQaType by ID
// @Description Get a BizQaType entry by ID
// @Tags biz_qa_type
// @Produce json
// @Param id path uint true "BizQaType ID"
// @Success 200 {object} models.BizQaType
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/biz_qa_type/{id} [get]
func (ctrl *BizQaTypeController) GetBizQaType(c *gin.Context) {
	id := c.Param("id")
	var qaType models.BizQaType
	if err := ctrl.DB.First(&qaType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, qaType)
}

func (ctrl *BizQaTypeController) UpdateBizQaType(c *gin.Context) {
	id := c.Param("id")
	var qaType models.BizQaType
	if err := ctrl.DB.First(&qaType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	if err := c.ShouldBindJSON(&qaType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.DB.Save(&qaType)
	c.JSON(http.StatusOK, qaType)
}

func (ctrl *BizQaTypeController) DeleteBizQaType(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.DB.Delete(&models.BizQaType{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted"})
}

// GetAllBizQaTypes godoc
// @Summary Get all BizQaTypes
// @Description Get all BizQaType entries
// @Tags biz_qa_type
// @Produce json
// @Success 200 {array} models.BizQaType
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/biz_qa_types [get]
func (ctrl *BizQaTypeController) GetAllBizQaTypes(c *gin.Context) {
	var qaTypes []models.BizQaType
	if err := ctrl.DB.Find(&qaTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to fetch records"})
		return
	}
	c.JSON(http.StatusOK, qaTypes)
}

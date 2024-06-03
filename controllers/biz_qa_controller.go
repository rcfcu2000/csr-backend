package controllers

import (
	"csr-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BizQaController struct {
	DB *gorm.DB
}

type CreateBizQaRequest struct {
	BizQa            models.BizQa             `json:"biz_qa"`
	BizQaQuestions   []models.BizQaQuestion   `json:"biz_qa_questions"`
	BizQuestionTypes []models.BizQuestionType `json:"biz_question_types"`
}

func NewBizQaController(db *gorm.DB) *BizQaController {
	return &BizQaController{DB: db}
}

func (ctrl *BizQaController) CreateBizQa_1(c *gin.Context) {
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

// CreateBizQaComplex godoc
// @Summary Create a BizQa with related entities
// @Description Create a BizQa along with BizQaQuestions and BizQuestionTypes in one transaction
// @Tags biz_qa
// @Accept json
// @Produce json
// @Param data body CreateBizQaRequest true "BizQa Complex Request"
// @Success 200 {object} models.BizQa
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/biz_qa_complex [post]
func (ctrl *BizQaController) CreateBizQaComplex(c *gin.Context) {
	var req CreateBizQaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	err := ctrl.DB.Transaction(func(tx *gorm.DB) error {
		req.BizQa.UpdateTime = time.Now()
		if err := tx.Create(&req.BizQa).Error; err != nil {
			return err
		}

		for _, question := range req.BizQaQuestions {
			question.Qid = req.BizQa.ID
			question.UpdateTime = time.Now()
			if err := tx.Create(&question).Error; err != nil {
				return err
			}
		}

		for _, questionType := range req.BizQuestionTypes {
			questionType.Qid = req.BizQa.ID
			questionType.UpdateTime = time.Now()
			if err := tx.Create(&questionType).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create records"})
		return
	}

	c.JSON(http.StatusOK, req.BizQa)
}

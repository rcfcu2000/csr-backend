package biz

import (
	"net/http"
	"strconv"
	"time"

	"xtt/global"
	"xtt/utils"

	models "xtt/model/biz"
	"xtt/model/common/request"
	"xtt/model/common/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BizQaController struct {
}

type CreateBizQaRequest struct {
	BizQa            models.BizQa             `json:"biz_qa"`
	BizQaQuestions   []models.BizQaQuestion   `json:"biz_qa_questions"`
	BizQuestionTypes []models.BizQuestionType `json:"biz_question_types"`
}

func (ctrl *BizQaController) CreateBizQa_1(c *gin.Context) {
	var newQa models.BizQa
	if err := c.ShouldBindJSON(&newQa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newQa.UpdateTime = time.Now()
	global.GVA_DB.Create(&newQa)
	c.JSON(http.StatusOK, newQa)
}

// GetBizQa handles fetching a qa by ID
// @Summary Get a merchant by ID
// @Description Get a merchant by ID
// @Tags biz_qa
// @Produce  json
// @Param id path string true "Qa ID"
// @Success 200 {object} models.BizQa
// @Failure 404 {object} models.ErrorResponse
// @Router /qa/get/{id} [get]
func (ctrl *BizQaController) GetBizQa(c *gin.Context) {
	id := c.Param("id")
	var qa models.BizQa
	if err := global.GVA_DB.First(&qa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, qa)
}

// GetBizQaByQuestion handles fetching qa by question
// @Summary Get qa by question
// @Description Get qa by question
// @Tags biz_qa
// @Produce  json
// @Param shopid path int true "Shopid"
// @Param questions body models.BizQaQuestions true "Qa"
// @Success 200 {object} []models.BizQa
// @Failure 404 {object} models.ErrorResponse
// @Router /qa/question/{shopid} [post]
func (ctrl *BizQaController) GetBizQaByQuestion(c *gin.Context) {
	id := c.Param("shopid")
	var ques models.BizQaQuestions

	if err := c.ShouldBindJSON(&ques); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sid, _ := strconv.Atoi(id)
	qa, err := qaService.GetBizQaByQuestion(ques.Questions, sid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, qa)
}

// UpdateBizQa handles updating an existing qa
// @Summary Update an existing qa
// @Description Update an existing qa
// @Tags biz_qa
// @Accept  json
// @Produce  json
// @Param id path string true "qa ID"
// @Param merchant body models.BizQa true "Qa"
// @Success 200 {object} models.BizQa
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /qa/update/{id} [put]
func (ctrl *BizQaController) UpdateBizQa(c *gin.Context) {
	id := c.Param("id")
	var qa models.BizQa
	if err := global.GVA_DB.First(&qa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&qa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := qaService.UpdateBizQa(&qa); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record update error"})
		return
	}
	c.JSON(http.StatusOK, qa)
}

// DeleteBizQa handles deleting a qa by ID
// @Summary Delete a qa by ID
// @Description Delete a qa by ID
// @Tags biz_qa
// @Produce  json
// @Param id path string true "Qa ID"
// @Success 200 {string} string "Qa deleted successfully"
// @Failure 500 {object} models.ErrorResponse
// @Router /qa/delete/{id} [delete]
func (ctrl *BizQaController) DeleteBizQa(c *gin.Context) {
	id := c.Param("id")

	qid, _ := strconv.Atoi(id)
	if err := qaService.DeleteBizQa(uint(qid)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found in qa question table"})
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
// @Router /qa/biz_qa_complex [post]
func (ctrl *BizQaController) CreateBizQaComplex(c *gin.Context) {
	var req CreateBizQaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	// err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	// 	req.BizQa.UpdateTime = time.Now()
	// 	if err := tx.Create(&req.BizQa).Error; err != nil {
	// 		return err
	// 	}

	// 	for _, qatype := range req.BizQa.QaTypes {
	// 		qaTypeService.IncrementRefCount(qatype.ID)
	// 	}

	// 	return nil
	// })

	req.BizQa.UpdateTime = time.Now()
	if err := global.GVA_DB.Create(&req.BizQa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, qatype := range req.BizQa.QaTypes {
		err := qaTypeService.IncrementRefCount(qatype.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create records"})
			return
		}
	}

	c.JSON(http.StatusOK, req.BizQa)
}

// GetQaList
// @Tags	  biz_qa
// @Summary   分页获取知识库列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo           true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取知识库列表,返回包括列表,总数,页码,每页数量"
// @Router    /qa/getQaList [post]
func (ctrl *BizQaController) GetQaList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	ktype, _ := strconv.Atoi(pageInfo.Keyword)
	list, total, err := qaService.GetQaList(pageInfo, ktype)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

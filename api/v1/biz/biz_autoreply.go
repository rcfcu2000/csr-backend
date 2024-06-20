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

type BizAutoReplyController struct {
}

func (ctrl *BizAutoReplyController) CreateAutoReply(c *gin.Context) {
	var newAr models.BizAutoReply
	if err := c.ShouldBindJSON(&newAr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newAr.UpdateTime = time.Now()
	global.GVA_DB.Create(&newAr)
	c.JSON(http.StatusOK, newAr)
}

// GetBizQa handles fetching auto reply by ID
// @Summary Get auto reply by ID
// @Description Get auto reply by ID
// @Tags biz_autoreply
// @Produce  json
// @Param id path string true "ar ID"
// @Success 200 {object} models.BizAutoReply
// @Failure 404 {object} models.ErrorResponse
// @Router /autoreply/get/{id} [get]
func (ctrl *BizAutoReplyController) GetAutoReply(c *gin.Context) {
	id := c.Param("id")
	var qa models.BizAutoReply
	if err := global.GVA_DB.First(&qa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, qa)
}

// GetBizQaByQuestion handles fetching qa by question
// @Summary Get qa by question
// @Description Get qa by question
// @Tags biz_autoreply
// @Produce  json
// @Param shopid path int true "Shopid"
// @Param questions body models.BizQaQuestions true "Qa"
// @Success 200 {object} []models.BizQa
// @Failure 404 {object} models.ErrorResponse
// @Router /autoreply/question/{shopid} [post]
func (ctrl *BizAutoReplyController) GetBizQaByQuestion(c *gin.Context) {
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
// @Tags biz_autoreply
// @Accept  json
// @Produce  json
// @Param id path string true "qa ID"
// @Param merchant body models.BizAutoReply true "Qa"
// @Success 200 {object} models.BizAutoReply
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /autoreply/update/{id} [put]
func (ctrl *BizAutoReplyController) UpdateAutoReply(c *gin.Context) {
	id := c.Param("id")
	var qa models.BizAutoReply
	if err := global.GVA_DB.First(&qa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&qa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := arService.UpdateAutoReply(&qa); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record update error"})
		return
	}
	c.JSON(http.StatusOK, qa)
}

// DeleteBizQa handles deleting a qa by ID
// @Summary Delete a qa by ID
// @Description Delete a qa by ID
// @Tags biz_autoreply
// @Produce  json
// @Param id path string true "Qa ID"
// @Success 200 {string} string "Qa deleted successfully"
// @Failure 500 {object} models.ErrorResponse
// @Router /autoreply/delete/{id} [delete]
func (ctrl *BizAutoReplyController) DeleteAutoReply(c *gin.Context) {
	id := c.Param("id")

	qid, _ := strconv.Atoi(id)
	if err := qaService.DeleteBizQa(uint(qid)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found in qa question table"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted"})
}

// GetList
// @Tags	  biz_autoreply
// @Summary   分页获取知识库列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo           true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取知识库列表,返回包括列表,总数,页码,每页数量"
// @Router    /autoreply/getList [post]
func (ctrl *BizAutoReplyController) GetAutoReplyList(c *gin.Context) {
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

	list, total, err := arService.GetAutoReplyList(pageInfo)
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

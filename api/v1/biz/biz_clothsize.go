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

type BizClothSizeController struct {
}

func (ctrl *BizClothSizeController) CreateClothSize(c *gin.Context) {
	var newAr models.BizClothSize
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
// @Tags biz_clothsize
// @Produce  json
// @Param id path string true "ar ID"
// @Success 200 {object} models.BizClothSize
// @Failure 404 {object} models.ErrorResponse
// @Router /clothsize/get/{id} [get]
func (ctrl *BizClothSizeController) GetClothSize(c *gin.Context) {
	id := c.Param("id")
	var qa models.BizClothSize
	if err := global.GVA_DB.First(&qa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, qa)
}

// GetBizQaByQuestion handles fetching qa by question
// @Summary Get qa by question
// @Description Get qa by question
// @Tags biz_clothsize
// @Produce  json
// @Param shopid path int true "Shopid"
// @Param questions body models.BizQaQuestions true "Qa"
// @Success 200 {object} []models.BizClothSize
// @Failure 404 {object} models.ErrorResponse
// @Router /clothsize/question/{shopid} [post]
func (ctrl *BizClothSizeController) GetBizQaByQuestion(c *gin.Context) {
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

// UpdateBizQa handles updating an existing cloth size
// @Summary Update an existing cloth size by ID
// @Description Update an existing cloth size by ID
// @Tags biz_clothsize
// @Accept  json
// @Produce  json
// @Param id path string true "cloth size by ID"
// @Param merchant body models.BizClothSize true "cloth size"
// @Success 200 {object} models.BizClothSize
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /clothsize/update/{id} [put]
func (ctrl *BizClothSizeController) UpdateClothSize(c *gin.Context) {
	id := c.Param("id")
	var qa models.BizClothSize
	if err := global.GVA_DB.First(&qa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if err := c.ShouldBindJSON(&qa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := csService.UpdateClothSizeInfo(&qa); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record update error"})
		return
	}
	c.JSON(http.StatusOK, qa)
}

// DeleteBizQa handles deleting a cloth size by ID
// @Summary Delete a cloth size by ID
// @Description Delete a cloth size by ID
// @Tags biz_clothsize
// @Produce  json
// @Param id path string true "cloth size ID"
// @Success 200 {string} string "cloth size deleted successfully"
// @Failure 500 {object} models.ErrorResponse
// @Router /clothsize/delete/{id} [delete]
func (ctrl *BizClothSizeController) DeleteClothSize(c *gin.Context) {
	id := c.Param("id")

	qid, _ := strconv.Atoi(id)
	if err := qaService.DeleteBizQa(uint(qid)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found in qa question table"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted"})
}

// GetList
// @Tags	  biz_clothsize
// @Summary   分页获取尺码表列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo           true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取尺码表列表,返回包括列表,总数,页码,每页数量"
// @Router    /clothsize/getList [post]
func (ctrl *BizClothSizeController) GetClothSizeList(c *gin.Context) {
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

	list, total, err := csService.GetClothSizeInfoList(pageInfo)
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
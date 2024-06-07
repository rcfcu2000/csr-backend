package biz

import (
	"net/http"
	"strconv"

	"xtt/global"
	"xtt/model/common/request"
	"xtt/model/common/response"
	"xtt/utils"

	models "xtt/model/biz"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BizQaTypeController struct {
}

func NewBizQaTypeController() *BizQaTypeController {
	return &BizQaTypeController{}
}

func (ctrl *BizQaTypeController) CreateBizQaType(c *gin.Context) {
	var newQaType models.BizQaType
	if err := c.ShouldBindJSON(&newQaType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newQaType.RefCount = 0
	global.GVA_DB.Create(&newQaType)
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
// @Router /qatype/biz_qa_type/{id} [get]
func (ctrl *BizQaTypeController) GetBizQaType(c *gin.Context) {
	id := c.Param("id")
	var qaType models.BizQaType
	if err := global.GVA_DB.First(&qaType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, qaType)
}

func (ctrl *BizQaTypeController) UpdateBizQaType(c *gin.Context) {
	id := c.Param("id")
	var qaType models.BizQaType
	if err := global.GVA_DB.First(&qaType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	if err := c.ShouldBindJSON(&qaType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	global.GVA_DB.Save(&qaType)
	c.JSON(http.StatusOK, qaType)
}

func (ctrl *BizQaTypeController) DeleteBizQaType(c *gin.Context) {
	id := c.Param("id")
	if err := global.GVA_DB.Delete(&models.BizQaType{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted"})
}

// GetAllBizQaTypes godoc
// @Summary Get all BizQaTypes
// @Description Get all BizQaType entries
// @Tags biz_qa_type
// @Param     data  body      request.PageInfo     true  "页码, 每页大小"
// @Produce json
// @Success 200 {array} models.BizQaType
// @Failure 500 {object} models.ErrorResponse
// @Router /qatype/biz_qa_types [post]
func (ctrl *BizQaTypeController) GetAllBizQaTypes(c *gin.Context) {
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
	list, total, err := qaTypeService.GetQaTypeList(pageInfo, ktype)
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

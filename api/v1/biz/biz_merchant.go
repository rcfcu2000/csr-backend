package biz

import (
	"net/http"
	"strconv"

	"xtt/global"
	models "xtt/model/biz"
	"xtt/model/common/request"
	"xtt/model/common/response"
	"xtt/utils"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BizMerchantController struct {
	DB *gorm.DB
}

// CreateMerchant handles the creation of a new merchant
// @Summary Create a new merchant
// @Description Create a new merchant
// @Tags Merchants
// @Accept  json
// @Produce  json
// @Param merchant body models.BizMerchant true "Merchant"
// @Success 200 {object} models.BizMerchant
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /merchant/create [post]
func (ctrl *BizMerchantController) CreateMerchant(c *gin.Context) {
	var merchant models.BizMerchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := merchantService.CreateMerchant(&merchant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": merchant})
}

// GetMerchant handles fetching a merchant by ID
// @Summary Get a merchant by ID
// @Description Get a merchant by ID
// @Tags Merchants
// @Produce  json
// @Param id path string true "Merchant ID"
// @Success 200 {object} models.BizMerchant
// @Failure 404 {object} models.ErrorResponse
// @Router /merchant/get/{id} [get]
func (ctrl *BizMerchantController) GetMerchant(c *gin.Context) {
	id := c.Param("id")
	merchant, err := merchantService.GetMerchant(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": merchant})
}

// GetMerchantList
// @Tags	  Merchants
// @Summary   分页获取商品列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取商品列表,返回包括列表,总数,页码,每页数量"
// @Router    /merchant/getMerchantList [post]
func (ctrl *BizMerchantController) GetMerchantList(c *gin.Context) {
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
	list, total, err := merchantService.GetMerchantList(pageInfo)
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

// UpdateMerchant handles updating an existing merchant
// @Summary Update an existing merchant
// @Description Update an existing merchant
// @Tags Merchants
// @Accept  json
// @Produce  json
// @Param id path string true "Merchant ID"
// @Param merchant body models.BizMerchant true "Merchant"
// @Success 200 {object} models.BizMerchant
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /merchant/update/{id} [put]
func (ctrl *BizMerchantController) UpdateMerchant(c *gin.Context) {
	id := c.Param("id")
	var merchant models.BizMerchant
	if err := c.ShouldBindJSON(&merchant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := merchantService.UpdateMerchant(id, &merchant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": merchant})
}

// DeleteMerchant handles deleting a merchant by ID
// @Summary Delete a merchant by ID
// @Description Delete a merchant by ID
// @Tags Merchants
// @Produce  json
// @Param id path string true "Merchant ID"
// @Success 200 {string} string "Merchant deleted successfully"
// @Failure 500 {object} models.ErrorResponse
// @Router /merchant/delete/{id} [delete]
func (ctrl *BizMerchantController) DeleteMerchant(c *gin.Context) {
	id := c.Param("id")
	if err := merchantService.DeleteMerchant(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Merchant deleted successfully"})
}

// UploadExcel handles the Excel file upload and data extraction
// @Summary Upload an Excel file
// @Description Upload an Excel file and extract data into the merchant table
// @Tags Merchants
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Excel file"
// @Success 200 {object} string "Merchants uploaded successfully"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /merchant/upload [post]
func (ctrl *BizMerchantController) UploadExcel(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer f.Close()

	excelFile, err := excelize.OpenReader(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Assuming the data is in the first sheet
	sheetName := excelFile.GetSheetName(0)
	rows, err := excelFile.GetRows(sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var merchants []models.BizMerchant

	for _, row := range rows[1:] { // Skip header row
		if len(row) < 3 {
			continue // Skip rows that don't have enough columns
		}

		var links []models.BizLinks
		id, _ := strconv.Atoi(row[1])
		link := models.BizLinks{
			TaobaoId:  uint(id),
			UpdatedBy: "system",
		}
		links = append(links, link)

		merchant := models.BizMerchant{
			Name:          row[0],
			Alias:         row[0],
			Information:   row[2],
			UpdatedBy:     "system",
			MerchantLinks: links,
		}

		merchants = append(merchants, merchant)
	}

	if err := merchantService.CreateMerchants(merchants); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Merchants uploaded successfully"})
}

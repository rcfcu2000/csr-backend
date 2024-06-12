package biz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"xtt/model/biz"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type ShopController struct {
}

// CreateShop godoc
// @Summary Create a shop
// @Description Create a new shop
// @Tags Shop
// @Accept json
// @Produce json
// @Param data body models.BizShop true "Create Shop"
// @Success 200 {object} models.BizShop
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /shop/create [post]
func (ctrl *ShopController) CreateShop(c *gin.Context) {
	var req models.BizShop
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := shopService.GetOrCreateCategory(req.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := shopService.CreateShop(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.BrandInfo = fillBrandInfo(&req)
	c.JSON(http.StatusOK, req)

	//go calculateBrandInfo(&req)
}

// GetShop godoc
// @Summary Get a shop by Name
// @Description Get details of a shop by name
// @Tags Shop
// @Produce json
// @Param name path string true "Shop Name"
// @Success 200 {object} models.BizShop
// @Failure 404 {object} models.ErrorResponse
// @Router /shop/get/{name} [get]
func (ctrl *ShopController) GetShop(c *gin.Context) {
	name := c.Param("name")
	shop, err := shopService.GetShopByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shop)
}

// GetShop godoc
// @Summary Get a shop by Name
// @Description Get details of a shop by name
// @Tags Shop
// @Produce json
// @Param id path string true "Shop ID"
// @Success 200 {object} models.BizShop
// @Failure 404 {object} models.ErrorResponse
// @Router /shop/getbyid/{id} [get]
func (ctrl *ShopController) GetShopByID(c *gin.Context) {
	id := c.Param("id")
	sid, _ := strconv.Atoi(id)
	shop, err := shopService.GetShopByID(uint(sid))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shop)
}

// UpdateShop godoc
// @Summary Update an existing shop
// @Description Update details of an existing shop
// @Tags Shop
// @Accept json
// @Produce json
// @Param id path string true "Shop ID"
// @Param data body models.BizShop true "Update Shop Request"
// @Success 200 {object} models.BizShop
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /shop/update/{id} [put]
func (ctrl *ShopController) UpdateShop(c *gin.Context) {
	id := c.Param("id")
	var req models.BizShop
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cid, _ := strconv.Atoi(id)

	shop, err := shopService.GetShopByID(uint(cid))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shop.BrandInfo = calculateBrandInfo(&req)
	//shop.UpdateTime = time.Now()

	if err := shopService.UpdateShop(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, req)
}

// DeleteShop godoc
// @Summary Delete a shop by ID
// @Description Delete a shop by ID
// @Tags Shop
// @Produce json
// @Param id path string true "Shop ID"
// @Success 200 {string} string "Shop deleted successfully"
// @Failure 500 {object} models.ErrorResponse
// @Router /shop/delete/{id} [delete]
func (ctrl *ShopController) DeleteShop(c *gin.Context) {
	id := c.Param("id")
	shopID, _ := strconv.Atoi(id)
	if err := shopService.DeleteShop(uint(shopID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shop deleted successfully"})
}

// ListShops godoc
// @Summary List all shops
// @Description Get a list of all shops
// @Tags Shop
// @Produce json
// @Success 200 {array} models.BizShop
// @Failure 500 {object} models.ErrorResponse
// @Router /shop/list [get]
func (ctrl *ShopController) ListShops(c *gin.Context) {
	shops, err := shopService.ListShops()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shops)
}

// ListCategories godoc
// @Summary List all categories
// @Description Get a list of all categories
// @Tags Shop
// @Produce json
// @Success 200 {array} models.BizCategory
// @Failure 500 {object} models.ErrorResponse
// @Router /shop/category_list [get]
func (ctrl *ShopController) ListCategories(c *gin.Context) {
	categories, err := shopService.ListCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// RequestPayload represents the JSON data to be sent in the POST request
type RequestPayload struct {
	Name       string `json:"brand_name"`
	Philosophy string `json:"brand_philosophy"`
	Advantages string `json:"brand_advantages"`
	Category   string `json:"main_product_categories"`
}

func fillBrandInfo(shop *models.BizShop) string {
	brandInfo := "我们的店铺昵称是" + shop.Nickname
	brandInfo += "我们的经营项目是" + shop.BrandManagement
	brandInfo += "我们的品牌理念是" + shop.BrandBelief
	brandInfo += "我们的核心卖点是" + shop.BrandAdvantage

	return brandInfo
}

func calculateBrandInfo(shop *models.BizShop) string {
	url := "https://www.zhihuige.cc/csrnew/api/generate_brand_info"
	fmt.Println("URL:>", url)

	// Marshal the request payload into JSON
	data := RequestPayload{
		Name:       shop.Name,
		Philosophy: shop.BrandBelief,
		Advantages: shop.BrandAdvantage,
		Category:   shop.Category.Name,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	type info struct {
		Response string `json:"response"`
	}
	var shop_info info
	json.Unmarshal(body, &shop_info)

	fmt.Println("response Body:", shop_info.Response)

	shop.BrandInfo = shop_info.Response
	shopService.UpdateShop(shop)
	return "updated brand info"
}

// UploadCategory handles the Excel file upload and data extraction
// @Summary Upload an Excel file
// @Description Upload an Excel file and extract data into the category table
// @Tags Merchants
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Excel file"
// @Success 200 {object} string "category uploaded successfully"
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /shop/upload_category [post]
func (ctrl *ShopController) UploadCategory(c *gin.Context) {
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

	for _, row := range rows[1:] { // Skip header row
		if len(row) < 3 {
			continue // Skip rows that don't have enough columns
		}

		category := models.BizCategory{
			Name:       row[0],
			PresetText: row[2],
			Level:      1,
		}
		if err := shopService.GetOrCreateCategory(category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Merchants uploaded successfully"})
}

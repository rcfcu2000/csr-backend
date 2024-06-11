package biz

import (
	"net/http"
	"strconv"

	"xtt/model/biz"

	"github.com/gin-gonic/gin"
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
// @Router /shop [post]
func (ctrl *ShopController) CreateShop(c *gin.Context) {
	var req models.BizShop
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.BrandInfo = calculateBrandInfo(&req)

	if err := shopService.CreateShop(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, req)
}

// GetShop godoc
// @Summary Get a shop by ID
// @Description Get details of a shop by ID
// @Tags Shop
// @Produce json
// @Param id path string true "Shop ID"
// @Success 200 {object} models.BizShop
// @Failure 404 {object} models.ErrorResponse
// @Router /shop/{id} [get]
func (ctrl *ShopController) GetShop(c *gin.Context) {
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
// @Router /shop/{id} [put]
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
// @Router /shop/{id} [delete]
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
// @Router /shop [get]
func (ctrl *ShopController) ListShops(c *gin.Context) {
	shops, err := shopService.ListShops()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shops)
}

func calculateBrandInfo(shop *models.BizShop) string {
	return "品牌理念: " + shop.BrandBelief + "; 品牌优势: " + shop.BrandAdvantage
}

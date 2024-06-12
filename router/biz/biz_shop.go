package biz

import (
	v1 "xtt/api/v1"

	"github.com/gin-gonic/gin"
)

type ShopRouter struct{}

func (e *ShopRouter) InitRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	priRouter := publicRouter.Group("shop")

	shopController := v1.ApiGroupApp.BizApiGroup.ShopController
	{
		priRouter.POST("create", shopController.CreateShop)
		priRouter.GET("get/:name", shopController.GetShop)
		priRouter.GET("getbyid/:id", shopController.GetShopByID)
		priRouter.PUT("update/:id", shopController.UpdateShop)
		priRouter.DELETE("delete/:id", shopController.DeleteShop)
		priRouter.GET("list", shopController.ListShops)
		priRouter.GET("category_list", shopController.ListCategories)
		priRouter.POST("upload_category", shopController.UploadCategory) // New route for Excel upload
	}
}

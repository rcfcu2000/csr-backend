package biz

import (
	v1 "xtt/api/v1"

	"github.com/gin-gonic/gin"
)

type ClothSizeInfoRouter struct{}

func (e *ClothSizeInfoRouter) InitRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	// priRouter := privateRouter.Group("inventory").Use(middleware.OperationRecord())
	priRouterWithoutRecord := publicRouter.Group("clothsize")

	csController := v1.ApiGroupApp.BizApiGroup.BizClothSizeController
	{
		priRouterWithoutRecord.POST("create", csController.CreateClothSize)
		priRouterWithoutRecord.GET("get/:id", csController.GetClothSize)
		priRouterWithoutRecord.POST("getList", csController.GetClothSizeList)
		priRouterWithoutRecord.PUT("update/:id", csController.UpdateClothSize)
		priRouterWithoutRecord.PUT("updateMerchantList", csController.UpdateMerchantList)
		priRouterWithoutRecord.GET("getMerchantList/:id", csController.GetMerchantList)
	}
}

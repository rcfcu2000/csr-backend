package biz

import (
	v1 "xtt/api/v1"

	"github.com/gin-gonic/gin"
)

type MerchantRouter struct{}

func (e *MerchantRouter) InitRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	// priRouter := privateRouter.Group("inventory").Use(middleware.OperationRecord())
	priRouter := publicRouter.Group("merchant")
	priRouterWithoutRecord := publicRouter.Group("merchant")

	merchantController := v1.ApiGroupApp.BizApiGroup.BizMerchantController
	{
		priRouter.POST("create", merchantController.CreateMerchant)
		priRouter.GET("get/:id", merchantController.GetMerchant)
		priRouter.GET("getByTid/:id", merchantController.GetMerchantsByTid)
		priRouter.PUT("update/:id", merchantController.UpdateMerchant)
		priRouter.DELETE("delete/:id", merchantController.DeleteMerchant)
		priRouter.POST("upload", merchantController.UploadExcel) // New route for Excel upload
	}

	{
		priRouterWithoutRecord.POST("getMerchantList", merchantController.GetMerchantList) // 分页获取商品列表
	}

}

package biz

import (
	v1 "xtt/api/v1"

	"github.com/gin-gonic/gin"
)

type QaRouter struct{}

func (e *QaRouter) InitRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	// priRouter := privateRouter.Group("inventory").Use(middleware.OperationRecord())
	priRouterWithoutRecord := publicRouter.Group("qa")

	qaController := v1.ApiGroupApp.BizApiGroup.BizQaController
	{
		priRouterWithoutRecord.POST("biz_qa_complex", qaController.CreateBizQaComplex)
		priRouterWithoutRecord.GET("biz_qa/:id", qaController.GetBizQa)
		priRouterWithoutRecord.POST("getQaList", qaController.GetQaList)
		priRouterWithoutRecord.PUT("biz_qa/:id", qaController.UpdateBizQa)
		priRouterWithoutRecord.DELETE("biz_qa/:id", qaController.DeleteBizQa)

	}

}

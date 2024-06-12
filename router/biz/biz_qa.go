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
		priRouterWithoutRecord.GET("get/:id", qaController.GetBizQa)
		priRouterWithoutRecord.POST("question/:shopid", qaController.GetBizQaByQuestion)
		priRouterWithoutRecord.POST("getQaList", qaController.GetQaList)
		priRouterWithoutRecord.PUT("update/:id", qaController.UpdateBizQa)
		priRouterWithoutRecord.DELETE("delete/:id", qaController.DeleteBizQa)

	}

}

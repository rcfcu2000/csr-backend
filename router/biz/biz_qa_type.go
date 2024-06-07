package biz

import (
	v1 "xtt/api/v1"

	"github.com/gin-gonic/gin"
)

type QaTypeRouter struct{}

func (e *QaTypeRouter) InitRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	// priRouter := privateRouter.Group("inventory").Use(middleware.OperationRecord())
	priRouterWithoutRecord := publicRouter.Group("qatype")

	qaTypeController := v1.ApiGroupApp.BizApiGroup.BizQaTypeController
	{
		priRouterWithoutRecord.POST("biz_qa_type", qaTypeController.CreateBizQaType)
		priRouterWithoutRecord.GET("biz_qa_type/:id", qaTypeController.GetBizQaType)
		priRouterWithoutRecord.POST("biz_qa_types", qaTypeController.GetAllBizQaTypes)
		priRouterWithoutRecord.PUT("biz_qa_type/:id", qaTypeController.UpdateBizQaType)
		priRouterWithoutRecord.DELETE("biz_qa_type/:id", qaTypeController.DeleteBizQaType)
	}

}

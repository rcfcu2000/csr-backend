package biz

import (
	v1 "xtt/api/v1"

	"github.com/gin-gonic/gin"
)

type AutoReplyRouter struct{}

func (e *AutoReplyRouter) InitRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	// priRouter := privateRouter.Group("inventory").Use(middleware.OperationRecord())
	priRouterWithoutRecord := publicRouter.Group("autoreply")

	arController := v1.ApiGroupApp.BizApiGroup.BizAutoReplyController
	{
		priRouterWithoutRecord.POST("create", arController.CreateAutoReply)
		priRouterWithoutRecord.GET("get/:id", arController.GetAutoReply)
		priRouterWithoutRecord.POST("getList", arController.GetAutoReplyList)
		priRouterWithoutRecord.PUT("update/:id", arController.UpdateAutoReply)

	}

}

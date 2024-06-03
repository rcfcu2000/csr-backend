package biz

import (
	v1 "xtt/api/v1"

	"github.com/gin-gonic/gin"
)

type MessageRouter struct{}

func (e *MessageRouter) InitRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	// priRouter := privateRouter.Group("inventory").Use(middleware.OperationRecord())
	priRouterWithoutRecord := publicRouter.Group("message")

	messageController := v1.ApiGroupApp.BizApiGroup.BizMessageController
	{
		priRouterWithoutRecord.POST("biz_messages", messageController.CreateBizMessages)
	}

}

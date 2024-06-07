package biz

import "xtt/service"

type ApiGroup struct {
	BizMessageController
	BizQaController
	BizQaTypeController
	BizMerchantController
}

var (
	merchantService = service.ServiceGroupApp.BizServiceGroup.MerchantService
)

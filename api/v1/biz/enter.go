package biz

import "xtt/service"

type ApiGroup struct {
	BizMessageController
	BizQaController
	BizQaTypeController
	BizMerchantController
}

var (
	qaService       = service.ServiceGroupApp.BizServiceGroup.BizQaService
	qaTypeService   = service.ServiceGroupApp.BizServiceGroup.BizQaTypeService
	merchantService = service.ServiceGroupApp.BizServiceGroup.MerchantService
)

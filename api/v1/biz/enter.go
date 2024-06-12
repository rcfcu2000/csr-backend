package biz

import "xtt/service"

type ApiGroup struct {
	BizMessageController
	BizQaController
	BizQaTypeController
	BizMerchantController
	ShopController
}

var (
	qaService       = service.ServiceGroupApp.BizServiceGroup.BizQaService
	qaTypeService   = service.ServiceGroupApp.BizServiceGroup.BizQaTypeService
	merchantService = service.ServiceGroupApp.BizServiceGroup.MerchantService
	shopService     = service.ServiceGroupApp.BizServiceGroup.BizShopService
)

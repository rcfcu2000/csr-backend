package biz

import "xtt/service"

type ApiGroup struct {
	BizMessageController
	BizQaController
	BizQaTypeController
	BizMerchantController
	ShopController
	BizAutoReplyController
	BizClothSizeController
}

var (
	qaService       = service.ServiceGroupApp.BizServiceGroup.BizQaService
	qaTypeService   = service.ServiceGroupApp.BizServiceGroup.BizQaTypeService
	merchantService = service.ServiceGroupApp.BizServiceGroup.MerchantService
	shopService     = service.ServiceGroupApp.BizServiceGroup.BizShopService
	arService       = service.ServiceGroupApp.BizServiceGroup.BizAutoReplyService
	csService       = service.ServiceGroupApp.BizServiceGroup.BizClothSizeService
)

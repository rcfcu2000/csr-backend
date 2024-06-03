package v1

import (
	"xtt/api/v1/biz"
	"xtt/api/v1/example"
	"xtt/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	BizApiGroup     biz.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

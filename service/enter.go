package service

import (
	"xtt/service/biz"
	"xtt/service/example"
	"xtt/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	BizServiceGroup biz.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

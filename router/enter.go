package router

import (
	"xtt/router/biz"
	"xtt/router/example"
	"xtt/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Biz     biz.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

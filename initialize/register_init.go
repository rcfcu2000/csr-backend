package initialize

import (
	_ "xtt/source/example"
	_ "xtt/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}

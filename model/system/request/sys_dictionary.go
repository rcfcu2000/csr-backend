package request

import (
	"xtt/model/common/request"
	"xtt/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}

package request

import (
	"xtt/model/common/request"
	"xtt/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}

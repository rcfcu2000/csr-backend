package request

import (
	"xtt/model/common/request"
	"xtt/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}

package request

import (
	"xtt/model/common/request"
	"xtt/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}

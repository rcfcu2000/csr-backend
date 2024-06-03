package response

import "xtt/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}

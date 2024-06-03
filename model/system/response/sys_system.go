package response

import "xtt/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}

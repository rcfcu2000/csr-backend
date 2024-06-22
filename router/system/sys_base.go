package system

import (
	v1 "xtt/api/v1"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("userinfo", baseApi.GetUserInfoByName) // 获取自身信息
		baseRouter.POST("ssoLogin", baseApi.UserSso)           // SSO 登录
		baseRouter.POST("rsaSsoLogin", baseApi.RSASso)         // RSA SSO 登录
		baseRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)     // 设置自身信息
	}
	return baseRouter
}

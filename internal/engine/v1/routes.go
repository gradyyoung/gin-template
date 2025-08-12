package v1

import (
	"github.com/gin-gonic/gin"
	"ygang.top/gin-template/internal/handler"
)

// Routes 路由
type Routes struct {
	SysUserHandler *handler.SysUserHandler
}

func NewRoutes(
	sysUserHandler *handler.SysUserHandler,
) *Routes {
	return &Routes{
		SysUserHandler: sysUserHandler,
	}
}

func (v Routes) SetupRoutes(engine *gin.Engine) {
	router := engine.Group("/v1")
	sysUser := router.Group("/sys_user")
	{
		sysUser.GET("/list", v.SysUserHandler.GetUserList)
	}
}

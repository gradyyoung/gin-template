package engine

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ygang.top/gin-template/internal/config"
	"ygang.top/gin-template/internal/engine/middleware"
)

// Router 路由管理接口
type Router interface {
	SetupRoutes(engine *gin.Engine)
}

// NewEngine 创建一个 gin.Engine，并启动服务
func NewEngine(
	routes Router,
	config *config.ApplicationConfig,
	errorHandleMiddleware *middleware.ErrorHandleMiddleware,
	authMiddleware *middleware.AuthMiddleware,
) *gin.Engine {
	gin.ForceConsoleColor()
	// 创建 gin.Engine
	engine := gin.Default()
	// 使用全局中间件
	{
		engine.Use(errorHandleMiddleware.Handler())
		engine.Use(authMiddleware.Handler())
	}
	routes.SetupRoutes(engine)
	engine.Run(fmt.Sprintf(":%d", config.Server.Port))
	return engine
}

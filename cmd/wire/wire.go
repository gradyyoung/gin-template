//go:build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"ygang.top/gin-template/internal/config"
	"ygang.top/gin-template/internal/database"
	"ygang.top/gin-template/internal/engine"
	"ygang.top/gin-template/internal/engine/api_v1"
	"ygang.top/gin-template/internal/engine/middleware"
	"ygang.top/gin-template/internal/handler"
	"ygang.top/gin-template/internal/service"
	"ygang.top/gin-template/util"
)

func InitApplication() *gin.Engine {
	wire.Build(
		config.InitApplicationConfig,
		database.ProviderSet,
		util.NewRedisClient,
		service.ProviderSet,
		handler.ProviderSet,
		middleware.ProviderSet,
		api_v1.NewRoutes,
		wire.Bind(new(engine.Router), new(*api_v1.Routes)), // 绑定接口
		engine.NewEngine,
	)
	return nil
}

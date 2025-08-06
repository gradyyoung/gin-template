package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"ygang.top/gin-template/internal/config"
	"ygang.top/gin-template/internal/dto"
	"ygang.top/gin-template/internal/errs"
	"ygang.top/gin-template/util"
)

type AuthMiddleware struct {
	Config      *config.ApplicationConfig
	RedisClient *util.RedisClient
}

func NewAuthMiddleware(
	config *config.ApplicationConfig,
	redisClient *util.RedisClient,
) *AuthMiddleware {
	return &AuthMiddleware{
		Config:      config,
		RedisClient: redisClient,
	}
}

func (a *AuthMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		excludeUrls := a.Config.Server.Auth.ExcludeUrls
		if excludeUrls != nil && len(excludeUrls) > 0 {
			for _, u := range excludeUrls {
				// 判断请求路径是否以 u 结尾
				if strings.HasSuffix(ctx.Request.URL.Path, u) {
					ctx.Next()
					return
				}
			}
		}
		token := ctx.GetHeader(a.Config.Server.Auth.Header)
		if token == "" {
			ctx.JSON(http.StatusOK, dto.FailedErrResponse(http.StatusNonAuthoritativeInfo, errs.UserNonAuthoritative))
			ctx.Abort()
			return
		}
		// 获取用户id
		var userId string
		// TODO: 校验 Token
		// 认证成功，设置用户id，并且续签token
		ctx.Set("userId", userId)
		// TODO: 续签
		ctx.Next()
	}
}

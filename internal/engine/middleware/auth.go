package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
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
	return func(c *gin.Context) {
		excludeUrls := a.Config.Server.Auth.ExcludeUrls
		if excludeUrls != nil && len(excludeUrls) > 0 {
			for _, u := range excludeUrls {
				// 判断请求路径是否以 u 结尾
				if strings.HasSuffix(c.Request.URL.Path, u) {
					c.Next()
					return
				}
			}
		}
		token := c.GetHeader(a.Config.Server.Auth.Header)
		if token != "" {
			userId, err := a.RedisClient.Get(util.LoginToken(token))
			if err != nil {
				logrus.Errorf("%+v\n", errors.Wrap(err, ""))
				c.JSON(http.StatusOK, dto.FailedMsgResponse(http.StatusInternalServerError, "服务器内部错误！"))
				c.Abort()
				return
			}
			if userId != "" {
				// 认证成功，设置用户id，并且续签token
				c.Set("userId", userId)
				a.RedisClient.Expire(util.LoginToken(token), time.Duration(a.Config.Server.Auth.TokenExpired)*time.Minute)
			} else {
				c.JSON(http.StatusOK, dto.FailedErrResponse(http.StatusNonAuthoritativeInfo, errs.UserNonAuthoritative))
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusOK, dto.FailedErrResponse(http.StatusNonAuthoritativeInfo, errs.UserNonAuthoritative))
			c.Abort()
			return
		}
		c.Next()
	}
}

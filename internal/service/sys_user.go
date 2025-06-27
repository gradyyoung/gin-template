package service

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
	"ygang.top/gin-template/internal/config"
	"ygang.top/gin-template/internal/dao"
	"ygang.top/gin-template/internal/errs"
	"ygang.top/gin-template/internal/model"
	"ygang.top/gin-template/util"
)

type SysUserService struct {
	Q           *dao.Query
	RedisClient *util.RedisClient
	Config      *config.ApplicationConfig
}

func NewSysUserService(
	Q *dao.Query,
	RedisClient *util.RedisClient,
	Config *config.ApplicationConfig,
) *SysUserService {
	return &SysUserService{
		Q:           Q,
		RedisClient: RedisClient,
		Config:      Config,
	}
}

// GetUserList 获取用户列表
func (s *SysUserService) GetUserList() ([]*model.SysUser, error) {
	find, err := s.Q.SysUser.Select().Find()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return find, nil
}

// UserLogin 用户登录
func (s *SysUserService) UserLogin(username, password string) (string, error) {
	user, err := s.Q.SysUser.Where(s.Q.SysUser.UserName.Eq(username)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.Wrap(errs.UserNameOrPasswordError, "")
		}
		return "", errors.Wrap(err, "")
	}
	if *user.Password != password {
		return "", errors.Wrap(errs.UserNameOrPasswordError, "")
	}
	// 创建 token
	token := uuid.New().String()
	err = s.RedisClient.SetEX(util.LoginToken(token), user.UserID, time.Duration(s.Config.Server.Auth.TokenExpired)*time.Minute)
	if err != nil {
		return "", errors.Wrap(err, "")
	}
	return token, nil
}

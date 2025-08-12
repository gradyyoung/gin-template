package service

import (
	"ygang.top/gin-template/internal/config"
	"ygang.top/gin-template/internal/dao"
	"ygang.top/gin-template/internal/model"
	"ygang.top/gin-template/internal/utils"
)

type SysUserService struct {
	SysUserDao  *dao.SysUserDao
	RedisClient *utils.RedisClient
	Config      *config.ApplicationConfig
}

func NewSysUserService(
	SysUserDao *dao.SysUserDao,
	RedisClient *utils.RedisClient,
	Config *config.ApplicationConfig,
) *SysUserService {
	return &SysUserService{
		SysUserDao:  SysUserDao,
		RedisClient: RedisClient,
		Config:      Config,
	}
}

// GetUserList 获取用户列表
func (s *SysUserService) GetUserList() ([]*model.SysUser, error) {
	list, err := s.SysUserDao.GetUserList()
	if err != nil {
		return nil, err
	}
	return list, nil
}

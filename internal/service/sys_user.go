package service

import (
	"ygang.top/gin-template/internal/dao"
	"ygang.top/gin-template/internal/model"
	"ygang.top/gin-template/util"
)

type SysUserService struct {
	Q           *dao.Query
	RedisClient *util.RedisClient
}

func NewSysUserService(
	Q *dao.Query,
	RedisClient *util.RedisClient,
) *SysUserService {
	return &SysUserService{
		Q:           Q,
		RedisClient: RedisClient,
	}
}

// GetUserList 获取用户列表
func (s *SysUserService) GetUserList() ([]*model.SysUser, error) {
	return s.Q.SysUser.Select().Find()
}

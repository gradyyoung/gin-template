package service

import (
	"ygang.top/gin-template/internal/dao"
	"ygang.top/gin-template/internal/model"
)

type SysUserService struct {
	Q *dao.Query
}

func NewSysUserService(Q *dao.Query) *SysUserService {
	return &SysUserService{
		Q: Q,
	}
}

// GetUserList 获取用户列表
func (s *SysUserService) GetUserList() ([]*model.SysUser, error) {
	return s.Q.SysUser.Select().Find()
}

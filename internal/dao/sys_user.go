package dao

import (
	"gorm.io/gorm"
	"ygang.top/gin-template/internal/model"
)

type SysUserDao struct {
	DB *gorm.DB
}

func NewSysUserDao(DB *gorm.DB) *SysUserDao {
	return &SysUserDao{
		DB: DB,
	}
}

func (d *SysUserDao) GetUserList() ([]*model.SysUser, error) {
	var list []*model.SysUser
	err := d.DB.Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

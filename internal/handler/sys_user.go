package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"ygang.top/gin-template/internal/dto"
	"ygang.top/gin-template/internal/service"
)

type SysUserHandler struct {
	SysUserService *service.SysUserService
}

func NewSysUserHandler(sysUserService *service.SysUserService) *SysUserHandler {
	return &SysUserHandler{
		SysUserService: sysUserService,
	}
}

// GetUserList 获取用户列表
func (h *SysUserHandler) GetUserList(ctx *gin.Context) {
	list, err := h.SysUserService.GetUserList()
	if err != nil {
		ctx.Error(errors.Wrap(err, ""))
	} else {
		ctx.JSON(http.StatusOK, dto.SuccessResponse(list, "获取用户列表成功"))
	}
}

// UserLogin 用户登录
func (h *SysUserHandler) UserLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	token, err := h.SysUserService.UserLogin(username, password)
	if err != nil {
		err := errors.Wrap(err, "")
		ctx.Error(err)
	} else {
		ctx.JSON(http.StatusOK, dto.SuccessResponse(token, "登录成功"))
	}
}

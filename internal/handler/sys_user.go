package handler

import (
	"github.com/gin-gonic/gin"
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
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse(err.Error()))
	} else {
		ctx.JSON(http.StatusOK, dto.SuccessResponse(list, "获取用户列表成功"))
	}

}

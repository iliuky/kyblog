package http

import (
	"kyblog/internal/blog/service"
	"kyblog/internal/common"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	srv    *service.Service
)

// InitHTTP 初始化控制器
func InitHTTP(s *service.Service) {
	logger = common.Application.Logger
	srv = s
}

// InitRouting 初始化路由
func InitRouting(route *gin.Engine) {
	route.GET("/", index)
	route.GET("/article/types", articleTypes)
}

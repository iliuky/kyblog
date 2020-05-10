package http

import (
	"kyblog/internal/blog/middleware"
	"kyblog/internal/blog/service"
	"kyblog/internal/common"
	"os"
	"path"
	"strings"

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

// InitResources 初始化静态资源
func InitResources(router *gin.Engine) {
	dir, _ := os.Getwd()
	suffix := "\\cmd\\blog"
	if strings.HasSuffix(dir, suffix) {
		dir = dir[:len(dir)-len(suffix)]
	}

	tmplExp := path.Join(dir, "templates", "blog", "*")
	publicExp := path.Join(dir, "public")
	router.LoadHTMLGlob(tmplExp)
	router.Static("/assets", publicExp)
}

// InitRouting 初始化路由
func InitRouting(router *gin.Engine) {
	g := router.Group("/", middleware.PageStatic(srv))
	{
		g.GET("/", articles)
		g.GET("/a/:pinyin", articleDetail)
		// g.GET("/article/types", articleTypes)
	}
}

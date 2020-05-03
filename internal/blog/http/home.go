package http

import "github.com/gin-gonic/gin"

// Index 首页
func index(context *gin.Context) {
	context.String(200, "blog")
}

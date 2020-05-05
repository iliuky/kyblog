package main

import (
	"kyblog/internal/blog/conf"
	"kyblog/internal/blog/http"
	"kyblog/internal/blog/service"
	"kyblog/internal/common"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := conf.Init()

	srv := service.NewService(config.ORM)
	srv.AutoMigrate()

	app := common.NewApp(config.BaseAppSettings)
	app.UseEndpoints(func(route *gin.Engine) {
		http.InitHTTP(srv)
		http.InitRouting(route)
		http.InitResources(route)
	})
	app.Start()
	app.Signal(func() {
		srv.Close()
	})
}

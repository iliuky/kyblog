package common

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Application 当前app
var Application *App

// Logger 当前 Logger
var Logger *zap.Logger

// App config
type App struct {
	// Name             string
	// Description      string
	Environment      string
	CurrentDirectory string
	Config           *BaseAppSettings
	Logger           *zap.Logger
	_engine          *gin.Engine
	_srv             *http.Server
}

// 路由配置委托
type engineFunc func(*gin.Engine)
type exitFunc func()

// NewApp 创建app
func NewApp(config *BaseAppSettings) *App {
	app := &App{
		Logger:      InitLogger("blog"),
		_engine:     initGin(),
		Config:      config,
		Environment: strings.ToLower(os.Getenv("ENVIRONMENT")),
	}
	app.CurrentDirectory, _ = os.Getwd()
	if app.Environment == "" {
		app.Environment = "development"
	}

	Application = app
	Logger = app.Logger
	return app
}

// Start 启动程序
func (app *App) Start() {
	app.Logger.Warn("启动程序")
	srv := &http.Server{
		Addr:    app.Config.HTTPServer.Addr,
		Handler: app._engine,
	}

	app._srv = srv
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal("http 侦听失败: ", zap.Error(err))
		}
	}()

	log.Printf("server url: http://%s", srv.Addr)
}

// Stop 程序停止
func (app *App) Stop() {
	app.Logger.Warn("程序停止中...")

	if app._srv != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := app._srv.Shutdown(ctx); err != nil {
			app.Logger.Fatal("http 停止失败")
		}
	}
	app.Logger.Warn("程序退出")
	app.Logger.Sync()
}

// UseEndpoints 使用站点路由配置
func (app *App) UseEndpoints(handle engineFunc) {
	handle(app._engine)
}

// Signal 侦听信号
func (app *App) Signal(handle exitFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-ch
		app.Logger.Warn("程序收到信号:", zap.String("signal", s.String()))

		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			app.Stop()
			handle()
			// time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

// IsDevelopment 是否开发环境
func (app *App) IsDevelopment() bool {
	return app.Environment == "development"
}

// IsProduction 是否生产环境
func (app *App) IsProduction() bool {
	return app.Environment == "production"
}

func initGin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	return engine
}

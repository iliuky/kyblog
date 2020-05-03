package conf

import (
	"kyblog/internal/common"
	"os"

	"github.com/prometheus/common/log"
)

// BlogSettings 站点配置
var BlogSettings *AppSettings

// AppSettings blog 配置文件
type AppSettings struct {
	*common.BaseAppSettings
}

// Init 初始化配置文件
func Init() *AppSettings {
	path := "appsettings.toml"
	environment := os.Getenv("ENVIRONMENT")

	if environment != "" {
		path = "appsettings." + environment + ".toml"
	}
	config := &AppSettings{&common.BaseAppSettings{}}
	err := config.Load(path)
	if err != nil {
		log.Fatal("初始化配置文件失败 %r", err)
	}
	BlogSettings = config
	return config
}

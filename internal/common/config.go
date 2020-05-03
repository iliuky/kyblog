package common

import (
	"kyblog/internal/orm"

	"github.com/BurntSushi/toml"
)

// ServerConfig http server config
type ServerConfig struct {
	Addr string
}

// BaseAppSettings http server config
type BaseAppSettings struct {
	HTTPServer *ServerConfig
	// orm
	ORM *orm.Config
}

// Load 加载配置文件
func (settings *BaseAppSettings) Load(path string) error {
	_, err := toml.DecodeFile(path, &settings)
	return err
}

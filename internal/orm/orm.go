package orm

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// Config mysql config.
type Config struct {
	DBType      string // 数据库类型
	DSN         string // 数据库连接
	Active      int    // pool
	Idle        int    // pool
	IdleTimeout int    // connect max life time.
	LogMode     bool   // 是否输出日记
}

// NewOrm 创建数据库gorm
func NewOrm(c *Config) (db *gorm.DB) {
	db, err := gorm.Open(c.DBType, c.DSN)
	if err != nil {
		log.Panic(c.DSN, err)
	}
	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout) * time.Minute)
	db.LogMode(c.LogMode)
	return
}

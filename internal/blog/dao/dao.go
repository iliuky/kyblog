package dao

import (
	"kyblog/internal/orm"

	"github.com/jinzhu/gorm"
)

// Dao 数据库访问
type Dao struct {
	DB *gorm.DB
}

// NewDao def
func NewDao(config *orm.Config) (d *Dao) {
	d = &Dao{
		DB: orm.NewOrm(config),
	}
	return
}

// Close def
func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
}

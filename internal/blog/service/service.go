package service

import (
	"kyblog/internal/blog/dao"
	"kyblog/internal/blog/model"
	"kyblog/internal/common"

	"github.com/jinzhu/gorm"
)

// Service def
type Service struct {
	dao *dao.Dao
	DB  *gorm.DB
}

// NewService 创建一个服务
func NewService(config *common.OrmConfig) (s *Service) {
	s = &Service{
		dao: dao.NewDao(config),
	}
	s.DB = s.dao.DB
	return s
}

// Close close all dao.
func (s *Service) Close() {
	s.dao.Close()
}

// AutoMigrate 自动迁移表
func (s *Service) AutoMigrate() {
	db := s.DB
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 auto_increment=1")
	db.AutoMigrate(&model.Article{}, &model.ArticleType{})
}

package service

import "kyblog/internal/blog/model"

// GetArticleAll 获取文章所有
func (s *Service) GetArticleAll() (list []*model.Article) {
	db := s.DB

	db.Where("status = ?", 0).Select("title, pin_yin, article_type, ctime").Order("ctime desc").Find(&list)
	return
}

// GetArticleTypeAll 获取文章分类
func (s *Service) GetArticleTypeAll() (list []*model.ArticleType) {
	db := s.DB

	db.Where("status = ?", 0).Order("sort desc").Find(&list)
	return
}

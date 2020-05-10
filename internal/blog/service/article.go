package service

import (
	"kyblog/internal/blog/model"
	"time"
)

// GetArticleAll 获取文章所有
func (s *Service) GetArticleAll() (list []*model.Article) {
	db := s.DB

	db.Where("status = ?", 0).Select("title, pin_yin, article_type, ctime").Order("ctime desc").Find(&list)
	return
}

// GetArticleNonStaticAll 获取尚未静态化文章
func (s *Service) GetArticleNonStaticAll() (list []*model.Article) {
	db := s.DB

	db.Where("status = ? and static_time <= modify_time", 0).Select("pin_yin").Find(&list)
	return
}

// UpdateArticleStaticTime 更新文章的静态化时间
func (s *Service) UpdateArticleStaticTime(pinYin string) {
	db := s.DB
	article := &model.Article{}
	db.First(article, "pin_yin = ?", pinYin)

	article.StaticTime = time.Now().Unix()
	db.Model(&article).Update(&article)
}

// GetArticle 获取文章详细
func (s *Service) GetArticle(pinYin string) (article *model.Article) {
	db := s.DB
	article = &model.Article{}
	db.First(article, "pin_yin = ?", pinYin)
	return
}

// GetArticleTypeAll 获取文章分类
func (s *Service) GetArticleTypeAll() (list []*model.ArticleType) {
	db := s.DB

	db.Where("status = ?", 0).Order("sort desc").Find(&list)
	return
}

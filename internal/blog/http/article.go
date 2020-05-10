package http

import (
	"html/template"
	"kyblog/internal/blog/viewmodel"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func articles(c *gin.Context) {
	list := srv.GetArticleAll()
	articlesMap := make(map[int32][]*viewmodel.ArticleListItemModel)
	for _, a := range list {
		arr := articlesMap[a.ArticleType]
		item := &viewmodel.ArticleListItemModel{
			Date:   time.Unix(a.CreateTime, 0).Format("2006-01-02"),
			PinYin: a.PinYin,
			Title:  a.Title,
		}
		arr = append(arr, item)
		articlesMap[a.ArticleType] = arr
	}

	var typeModels []viewmodel.ArticleListTypeModel
	if len(list) > 0 {
		types := srv.GetArticleTypeAll()
		for _, t := range types {
			v, ok := articlesMap[t.ID]
			if ok {
				item := viewmodel.ArticleListTypeModel{
					TypeName:   t.Name,
					CreateTime: t.CreateTime,
					Articles:   v,
				}
				typeModels = append(typeModels, item)
			}
		}
	}
	// c.JSON(200, typeModels)
	c.HTML(200, "article_list.html", &viewmodel.ArticleListModel{
		ArticleTypes: &typeModels,
	})
}

func articleDetail(c *gin.Context) {
	pinyin := c.Param("pinyin")
	article := srv.GetArticle(pinyin)

	var html template.HTML
	if article.ContentType == "md" {
		var buff []byte = []byte(article.Content)
		output := blackfriday.Run(buff)
		html = template.HTML(output)
	} else {
		html = template.HTML(article.Content)
	}

	c.HTML(200, "article_detail.html", &viewmodel.ArticleDetailModel{
		Article: *article,
		HTML:    html,
		Date:    time.Unix(article.CreateTime, 0).Format("2006-01-02"),
	})
}

func articleTypes(c *gin.Context) {
	c.JSON(200, srv.GetArticleTypeAll())
}

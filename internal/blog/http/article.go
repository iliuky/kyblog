package http

import (
	"kyblog/internal/blog/viewmodel"
	"time"

	"github.com/gin-gonic/gin"
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
	c.HTML(200, "article_list.html", viewmodel.ArticleListModel{
		ArticleTypes: &typeModels,
	})
}

func articleTypes(c *gin.Context) {
	c.JSON(200, srv.GetArticleTypeAll())
}

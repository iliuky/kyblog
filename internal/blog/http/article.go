package http

import "github.com/gin-gonic/gin"

func articles(context *gin.Context) {

}

func articleTypes(context *gin.Context) {
	context.JSON(200, srv.GetArticleTypeAll())
}

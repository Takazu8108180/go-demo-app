package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/takazu8108180/go-demo-app/handlers"
	"github.com/takazu8108180/go-demo-app/infra/database"
	"github.com/takazu8108180/go-demo-app/repositories"
)

func ArticleRouter(router *gin.Engine, db *database.DB) {
	// userRouter := router.Group("/users")
	{
		h := handlers.NewHandler(repositories.NewArticleRepository(db))
		router.GET("/", h.HelloHandler)
		router.POST("/article", h.PostArticleHandler)
		router.GET("/article/list", h.GetArticleListHandler)
		router.GET("/article/detail/:id", h.GetArticleDetailHandler)
		router.POST("/article/nice", h.PostNiceHandler)
		router.POST("/comment", h.PostCommentHandler)
	}
}

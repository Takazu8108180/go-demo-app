package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	h "github.com/takazu8108180/go-demo-app/handlers"
)

func main() {
	r := gin.Default()

	r.GET("/", h.HelloHandler)
	r.POST("/article", h.PostArticleHandler)
	r.GET("/article/list", h.GetArticleListHandler)
	r.GET("/article/detail/:id", h.GetArticleDetailHandler)
	r.POST("/article/nice", h.PostNiceHandler)
	r.POST("/comment", h.PostCommentHandler)

	log.Println("Server start at port 8080")

	r.Run(":8080")
}

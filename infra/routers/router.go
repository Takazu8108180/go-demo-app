package routers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/takazu8108180/go-demo-app/infra/database"
)

type Router struct {
	router *gin.Engine
	db     *database.DB
}

// NewRouter コンストラクタ
func NewRouter(router *gin.Engine, db *database.DB) *Router {
	return &Router{
		router: router,
		db:     db,
	}
}

func (r *Router) Run() error {
	log.Println("Server start at port 8080")

	ArticleRouter(r.router, r.db)

	return r.router.Run(":8080")
}

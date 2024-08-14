package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/takazu8108180/go-demo-app/models"
	"github.com/takazu8108180/go-demo-app/repositories"
)

type Handler struct {
	ar *repositories.ArticleRepository
}

func NewHandler(ar *repositories.ArticleRepository) *Handler {

	return &Handler{
		ar: ar,
	}
}

func (h *Handler) HelloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!\n")
}

func (h *Handler) PostArticleHandler(c *gin.Context) {

	var reqArticle models.Article

	if err := c.BindJSON(&reqArticle); err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, reqArticle)
}

func (h *Handler) GetArticleListHandler(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageNum, err := strconv.Atoi(page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid query parameter"})
		c.Abort()
		return
	}

	articles, err := h.ar.GetList(c, pageNum)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, articles)
}

func (h *Handler) GetArticleDetailHandler(c *gin.Context) {
	articleID := c.Param("id")

	log.Println(articleID)

	article, err := h.ar.GetByID(c, articleID)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, article)
}

func (h *Handler) PostNiceHandler(c *gin.Context) {

	var niceArticle models.Article

	if err := c.BindJSON(&niceArticle); err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, niceArticle)
}

func (h *Handler) PostCommentHandler(c *gin.Context) {

	var comment models.Comment

	if err := c.BindJSON(&comment); err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, comment)
}

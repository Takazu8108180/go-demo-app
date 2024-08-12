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

	log.Println(pageNum)

	// TODO:DB操作に置換
	// articleList := []models.Article{models.Article1, models.Article2}
	articleList, err := h.ar.GetList(c, pageNum)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, articleList)
}

func (h *Handler) GetArticleDetailHandler(c *gin.Context) {
	id := c.Param("id")
	articleID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid query parameter"})
		c.Abort()
		return
	}

	log.Println(articleID)

	// TODO:DB操作に置換
	article := models.Article1

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

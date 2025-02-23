package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/takazu8108180/go-demo-app/adapter/presenter/model"
	"github.com/takazu8108180/go-demo-app/models"
	"github.com/takazu8108180/go-demo-app/usecase"
)

type Handler struct {
	au *usecase.ArticleUseCase
}

func NewHandler(au *usecase.ArticleUseCase) *Handler {
	return &Handler{
		au: au,
	}
}

func (h *Handler) HelloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!\n")
}

func (h *Handler) PostArticleHandler(c *gin.Context) {
	var reqBody model.CreateArticleRequestBody
	if err := c.BindJSON(&reqBody); err != nil {
		return
	}

	article, err := h.au.CreateArticleUseCase(c, &reqBody)
	if err != nil {
		return
	}

	var resBody model.CreateArticleResponseBody
	resBody.ID = article.ID
	resBody.Title = article.Title
	resBody.Contents = article.Contents
	resBody.Username = article.Username
	resBody.CreatedAt = article.CreatedAt

	c.IndentedJSON(http.StatusOK, resBody)
}

func (h *Handler) GetArticleListHandler(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	// log.Println(page)
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid query parameter"})
		c.Abort()
		return
	}

	articles, err := h.au.GetArticleListuseCase(c, pageNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Internal Server Error"})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, articles)
}

func (h *Handler) GetArticleDetailHandler(c *gin.Context) {
	articleID := c.Param("id")
	log.Println(articleID)

	article, err := h.au.GetArticleByIDUseCase(c, articleID)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, article)
}

func (h *Handler) PostNiceHandler(c *gin.Context) {
	articleID := c.Param("id")
	article, err := h.au.UpdateNiceUseCase(c, articleID)
	if err != nil {
		return
	}

	var resBody model.SendNiceResponseBody
	resBody.ID = article.ID
	resBody.Title = article.Title
	resBody.Contents = article.Contents
	resBody.Username = article.Username
	resBody.CreatedAt = article.CreatedAt

	c.IndentedJSON(http.StatusOK, resBody)
}

func (h *Handler) PostCommentHandler(c *gin.Context) {

	var comment *models.Comment

	if err := c.BindJSON(&comment); err != nil {
		return
	}

	comment, err := h.au.CreateCommentUseCase(c, comment)
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, comment)
}

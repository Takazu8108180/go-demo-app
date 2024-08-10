package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/takazu8108180/go-demo-app/models"
)

func HelloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!\n")
}

func PostArticleHandler(c *gin.Context) {

	var reqArticle models.Article

	if err := c.BindJSON(&reqArticle); err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, reqArticle)
}

func GetArticleListHandler(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageNum, err := strconv.Atoi(page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid query parameter"})
		c.Abort()
		return
	}

	log.Println(pageNum)

	articleList := []models.Article{models.Article1, models.Article2}

	c.IndentedJSON(http.StatusOK, articleList)
}

func GetArticleDetailHandler(c *gin.Context) {
	id := c.Param("id")
	articleID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid query parameter"})
		c.Abort()
		return
	}

	log.Println(articleID)

	article := models.Article1

	c.IndentedJSON(http.StatusOK, article)
}

func PostNiceHandler(c *gin.Context) {

	var niceArticle models.Article

	if err := c.BindJSON(&niceArticle); err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, niceArticle)
}

func PostCommentHandler(c *gin.Context) {

	var comment models.Comment

	if err := c.BindJSON(&comment); err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, comment)
}

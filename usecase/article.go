package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/takazu8108180/go-demo-app/adapter/presenter/model"
	"github.com/takazu8108180/go-demo-app/models"
	"github.com/takazu8108180/go-demo-app/repositories"
)

type ArticleUseCase struct {
	ar *repositories.ArticleRepository
}

func NewArticleUseCase(ar *repositories.ArticleRepository) *ArticleUseCase {
	return &ArticleUseCase{
		ar: ar,
	}
}

func (au *ArticleUseCase) CreateArticleUseCase(c *gin.Context, reqBody *model.CreateArticleRequestBody) (*models.Article, error) {
	// TODO: reqBodyの型をadapterのモデルからusecaseのモデルに変換する

	// TODO: domain modelに変換
	var article models.Article
	article.Title = reqBody.Title
	article.Contents = reqBody.Contents
	article.Username = reqBody.Username

	output, err := au.ar.CreateArticle(c, &article)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (au *ArticleUseCase) GetArticleListuseCase(c *gin.Context, pageNum int) (*[]models.Article, error) {
	articles, err := au.ar.GetList(c, pageNum)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (au *ArticleUseCase) GetArticleByIDUseCase(c *gin.Context, id string) (*models.Article, error) {
	article, err := au.ar.GetByID(c, id)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (au *ArticleUseCase) UpdateNiceUseCase(c *gin.Context, articleID string) (*models.Article, error) {
	article, err := au.ar.UpdateNice(c, articleID)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (au *ArticleUseCase) CreateCommentUseCase(c *gin.Context, comment *models.Comment) (*models.Comment, error) {
	output, err := au.ar.CreateComment(c, comment)
	if err != nil {
		return nil, err
	}

	return output, nil
}

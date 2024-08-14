package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"

	"github.com/takazu8108180/go-demo-app/ent/article"
	"github.com/takazu8108180/go-demo-app/infra/database"
	"github.com/takazu8108180/go-demo-app/models"
)

type ArticleRepository struct {
	db *database.DB
}

func NewArticleRepository(db *database.DB) *ArticleRepository {
	return &ArticleRepository{
		db: db,
	}
}

func (ar *ArticleRepository) GetList(ctx context.Context, page int) (*[]models.Article, error) {
	result, err := ar.db.Client.Article.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	articles := make([]models.Article, len(result))

	for _, v := range result {
		var article models.Article

		article.ID = v.ID
		article.Title = v.Title
		article.Contents = v.Contents
		article.Username = v.Username
		article.NiceNum = v.Nice
		article.CreatedAt = v.CreatedAt

		articles = append(articles, article)
	}

	log.Println("user returned: ", articles)

	return &articles, nil
}

func (ar *ArticleRepository) GetByID(ctx context.Context, id string) (*models.Article, error) {
	articleID, _ := uuid.Parse(id)

	result, err := ar.db.Client.Article.
		Query().
		Where(article.ID(articleID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	var article models.Article

	article.ID = result[0].ID
	article.Title = result[0].Title
	article.Contents = result[0].Contents
	article.Username = result[0].Username
	article.NiceNum = result[0].Nice
	article.CreatedAt = result[0].CreatedAt

	log.Println("user returned: ", article)

	return &article, nil
}

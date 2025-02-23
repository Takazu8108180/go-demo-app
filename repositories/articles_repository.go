package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"

	"github.com/takazu8108180/go-demo-app/ent"

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
	articleID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	result, err := ar.db.Client.Article.Get(ctx, articleID)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	var article models.Article

	article.ID = result.ID
	article.Title = result.Title
	article.Contents = result.Contents
	article.Username = result.Username
	article.NiceNum = result.Nice
	article.CreatedAt = result.CreatedAt

	log.Println("user returned: ", article)

	return &article, nil
}

func (ar *ArticleRepository) CreateArticle(ctx context.Context, article *models.Article) (*models.Article, error) {
	//TODO: entでnewしている実装からドメイン層で生成するように改修する
	// id, err := uuid.NewRandom()

	result, err := ar.db.Client.Article.
		Create().
		SetTitle(article.Title).
		SetContents(article.Contents).
		SetUsername(article.Username).
		SetNice(1). //TODO: entのPositiveの設定を外す
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	var output models.Article

	output.ID = result.ID
	output.Title = result.Title
	output.Contents = result.Contents
	output.Username = result.Username
	output.CreatedAt = result.CreatedAt

	log.Println("user returned: ", output)

	return &output, nil
}

func (ar *ArticleRepository) UpdateNice(ctx context.Context, id string) (*models.Article, error) {
	articleID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	result, err := ar.db.Client.Article.
		UpdateOneID(articleID).
		AddNice(1).
		Save(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, fmt.Errorf("todo item was not found: %w", err)
	case err != nil:
		return nil, fmt.Errorf("update error: %w", err)
	}

	var output models.Article

	output.ID = result.ID
	output.Title = result.Title
	output.Contents = result.Contents
	output.Username = result.Username
	output.CreatedAt = result.CreatedAt

	log.Println("user returned: ", output)

	return &output, nil
}

func (ar *ArticleRepository) CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	result, err := ar.db.Client.Comment.
		Create().
		SetID(comment.CommentID).
		SetMessage(comment.Message).
		SetCreatedAt(comment.CreatedAt).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	var output models.Comment

	output.CommentID = result.ID
	output.Message = result.Message
	output.CreatedAt = result.CreatedAt

	log.Println("user returned: ", output)

	return &output, nil
}

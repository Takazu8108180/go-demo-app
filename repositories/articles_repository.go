package repositories

import (
	"context"

	"github.com/pkg/errors"

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

func (ar ArticleRepository) GetList(ctx context.Context, page int) (*[]models.Article, error) {
	// u, err := client.User.
	//     Query().
	//     Where(user.Name("a8m")).
	//     // `Only` fails if no user found,
	//     // or more than 1 user returned.
	//     Only(ctx)
	// if err != nil {
	//     return nil, fmt.Errorf("failed querying user: %w", err)
	// }
	// log.Println("user returned: ", u)
	// return u, nil

	return nil, errors.New("not implemented")
}

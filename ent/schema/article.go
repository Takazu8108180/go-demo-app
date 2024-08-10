package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("title").
			Default("unknown").
			MaxLen(100),
		field.String("contents").
			Default("unknown"),
		field.String("username").
			MaxLen(100).
			Default("unknown"),
		field.Int("nice").
			Positive(),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comment", Comment.Type),
	}
}

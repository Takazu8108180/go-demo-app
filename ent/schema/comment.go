package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("message").
			Default("unknown"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Movie struct {
	ent.Schema
}

func (Movie) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Int("release_year"),
		field.Enum("quality").
			Values("720p", "1080p", "4K"),
	}
}

func (Movie) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("casts", Cast.Type),
	}
}

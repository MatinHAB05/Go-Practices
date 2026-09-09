package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Cast struct {
	ent.Schema
}

func (Cast) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (Cast) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("movies", Movie.Type).
			Ref("casts"),
	}
}

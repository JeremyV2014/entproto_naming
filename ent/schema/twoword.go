package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// TwoWord holds the schema definition for the TwoWord entity.
type TwoWord struct {
	ent.Schema
}

// Fields of the TwoWord.
func (TwoWord) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("enums").
			Values("option1", "option2").
			Default("option1").
			Annotations(
				entproto.Field(2),
				entproto.Enum(map[string]int32{
					"option1": 0,
					"option2": 1,
				}),
			),
	}
}

// Edges of the TwoWord.
func (TwoWord) Edges() []ent.Edge {
	return nil
}

func (TwoWord) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}

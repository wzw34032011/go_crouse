package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("age").Positive(),
		field.String("name").Default("unknown"),
		field.Enum("gender").Values("0", "1"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

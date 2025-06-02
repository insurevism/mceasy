package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User schema
type User struct {
	ent.Schema
}

// Fields defines the User entity fields.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Unique().
			Immutable(),

		field.String("fullname").
			NotEmpty(),

		field.String("username").
			NotEmpty(),

		field.String("email").
			NotEmpty(),

		field.String("password").
			NotEmpty().
			Sensitive(),

		field.String("avatar").
			Optional(),
	}
}

// Mixin for shared fields
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseFieldMixin{},
	}
}

// Edges of the TransferOrder.
func (User) Edges() []ent.Edge {
	return nil
}

// Annotations allows adding extra metadata.
func (User) Annotations() []schema.Annotation {
	return nil
}

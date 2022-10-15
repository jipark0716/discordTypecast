package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UserTypecastSetting holds the schema definition for the UserTypecastSetting entity.
type UserTypecastSetting struct {
	ent.Schema
}

// Fields of the UserTypecastSetting.
func (UserTypecastSetting) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").Unique(),
		field.String("actor_id"),
		field.String("lang").
			Default("auto"),
		field.Int("max_seconds").
			Default(30),
		field.Float("naturalness").
			Default(0.8),
		field.Int("speed_x").
			Default(1),
		field.String("gid").
			Default("f8ON1ZpeF0mDFjZTQlr9G"),
		field.Int("style_idx").
			Default(0),
		field.String("last_pitch").
			Optional().
			Nillable(),
		field.String("mode").
			Optional().
			Nillable(),
		field.Int("pitch").
			Default(0),
		field.String("style_label").
			Default("0"),
		field.String("style_label_version").
			Default("v2"),
		field.Int("tempo").
			Default(0),
	}
}

// Edges of the UserTypecastSetting.
func (UserTypecastSetting) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TypecastMessage holds the schema definition for the TypecastMessage entity.
type TypecastMessage struct {
	ent.Schema
}

// Fields of the TypecastMessage.
func (TypecastMessage) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id"),
		field.String("text"),
		field.String("channel_id"),
		field.String("actor_id"),
		field.String("lang"),
		field.Int("max_seconds"),
		field.Float("naturalness"),
		field.Int("speed_x"),
		field.String("gid"),
		field.Int("style_idx"),
		field.String("last_pitch").
			Nillable(),
		field.String("mode").
			Nillable(),
		field.Int("pitch"),
		field.String("style_label"),
		field.String("style_label_version"),
		field.Int("tempo"),
		field.Int8("status").
			Default(0),
		field.Time("send_at").
			Nillable().
			Optional(),
	}
}

// Edges of the TypecastMessage.
func (TypecastMessage) Edges() []ent.Edge {
	return nil
}

func (TypecastMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

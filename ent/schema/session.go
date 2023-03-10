package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").
			Default(""),
		field.String("error").
			Default(""),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return nil
}

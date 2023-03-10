package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Session holds the schema definition for the Session entity.
type Dialog struct {
	ent.Schema
}

func (Dialog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("sid"),
		field.String("q").
			Default(""),
		field.String("a").
			Default(""),
		field.String("error").
			Default(""),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Session.
func (Dialog) Edges() []ent.Edge {
	return nil
}

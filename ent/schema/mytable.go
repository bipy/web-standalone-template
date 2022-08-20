package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// MyTable holds the schema definition for the MyTable entity.
type MyTable struct {
	ent.Schema
}

// Fields of the MyTable.
func (MyTable) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Positive().
			Immutable(),
		field.Text("name").
			NotEmpty(),
		field.Time("create_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the MyTable.
func (MyTable) Edges() []ent.Edge {
	return nil
}

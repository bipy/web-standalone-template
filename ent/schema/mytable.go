package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
		field.Text("value").
			NotEmpty(),
	}
}

// Edges of the MyTable.
func (MyTable) Edges() []ent.Edge {
	return nil
}

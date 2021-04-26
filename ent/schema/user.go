package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	player_pb "example/pkg/player/v1"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("Status").GoType(player_pb.Status_OK),
		field.JSON("Info", player_pb.Info{}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

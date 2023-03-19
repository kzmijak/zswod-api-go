package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type UuidMixin struct {
	mixin.Schema
}

func (UuidMixin) Fields() []ent.Field {
	return []ent.Field {
			field.UUID("id", uuid.UUID{}).
				Unique().
				Immutable().
				Default(uuid.New),
	}
}

type IconMixin struct {
	mixin.Schema
}

func (IconMixin) Fields() []ent.Field {
	return []ent.Field {
			field.String("iconId").Optional(),
	}
}
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// PhotographCity holds the schema definition for the PhotographCity entity.
type PhotographCity struct {
	ent.Schema
}

// Fields of the PhotographCity.
func (PhotographCity) Fields() []ent.Field {
	return []ent.Field{
		field.String("province_code").Nillable().Optional().Comment("省份id"),
		field.String("province_name").Nillable().Optional().Comment("省份名称"),
		field.String("city_code").Comment("城市id"),
		field.String("city_name").Comment("城市名称"),
	}
}

func (PhotographCity) Edges() []ent.Edge {
	return nil
}

// Annotations of the PhotographCity.
func (PhotographCity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "photograph_city"},
	}
}

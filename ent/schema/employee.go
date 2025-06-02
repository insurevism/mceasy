package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Unique().
			Immutable(),

		field.String("employee_id").
			MaxLen(20).
			Unique().
			Optional().
			Comment("Unique employee identifier like EMP-0001, EMP-0002"),

		field.String("full_name").
			MaxLen(255).
			NotEmpty(),

		field.String("email").
			MaxLen(255).
			NotEmpty().
			Unique(),

		field.String("phone").
			MaxLen(20).
			Optional(),

		field.String("position").
			MaxLen(100).
			Optional(),

		field.String("department").
			MaxLen(100).
			Optional(),

		field.Time("hire_date").
			Comment("Employee hire date"),

		field.Float("base_salary").
			Default(10000000.00).
			Comment("Base salary in IDR"),

		field.Bool("is_active").
			Default(true),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attendances", Attendance.Type),
		edge.To("salary_calculations", SalaryCalculation.Type),
	}
}

// Mixin for shared fields
func (Employee) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseFieldMixin{},
	}
}

// Indexes of the Employee.
func (Employee) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("employee_id"),
		index.Fields("email"),
		index.Fields("is_active"),
	}
}

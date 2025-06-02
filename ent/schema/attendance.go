package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Attendance holds the schema definition for the Attendance entity.
type Attendance struct {
	ent.Schema
}

// Fields of the Attendance.
func (Attendance) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Unique().
			Immutable(),

		field.Uint64("employee_id").
			Comment("Foreign key to employees table"),

		field.Time("attendance_date").
			Comment("Date of attendance"),

		field.Time("check_in_time").
			Optional().
			Comment("Actual check-in time"),

		field.Time("check_out_time").
			Optional().
			Comment("Actual check-out time"),

		field.Enum("status").
			Values("present", "absent", "late", "half_day").
			Default("absent"),

		field.Bool("is_weekend").
			Default(false).
			Comment("True for Saturday and Sunday"),

		field.Text("notes").
			Optional().
			Comment("Additional notes for attendance"),

		field.Bool("marked_by_admin").
			Default(false).
			Comment("True if manually marked by admin"),
	}
}

// Edges of the Attendance.
func (Attendance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("employee", Employee.Type).
			Ref("attendances").
			Field("employee_id").
			Unique().
			Required(),
	}
}

// Mixin for shared fields
func (Attendance) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseFieldMixin{},
	}
}

// Indexes of the Attendance.
func (Attendance) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("employee_id", "attendance_date").Unique(),
		index.Fields("attendance_date"),
		index.Fields("employee_id"),
		index.Fields("status"),
		index.Fields("is_weekend"),
	}
}

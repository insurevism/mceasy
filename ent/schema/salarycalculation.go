package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// SalaryCalculation holds the schema definition for the SalaryCalculation entity.
type SalaryCalculation struct {
	ent.Schema
}

// Fields of the SalaryCalculation.
func (SalaryCalculation) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Unique().
			Immutable(),

		field.Uint64("employee_id").
			Comment("Foreign key to employees table"),

		field.Time("calculation_month").
			Comment("First day of the month for calculation (YYYY-MM-01)"),

		field.Float("base_salary").
			Comment("Base salary for the month"),

		field.Int("total_working_days").
			Comment("Total working days in the month (excluding weekends)"),

		field.Int("absent_days").
			Default(0).
			Comment("Number of absent working days"),

		field.Int("present_days").
			Default(0).
			Comment("Number of present working days"),

		field.Float("final_salary").
			Comment("Final calculated salary after deductions"),

		field.Float("deduction_amount").
			Default(0.00).
			Comment("Total deduction amount"),

		field.Text("calculation_formula").
			Optional().
			Comment("Formula used for calculation (for audit purposes)"),
	}
}

// Edges of the SalaryCalculation.
func (SalaryCalculation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("employee", Employee.Type).
			Ref("salary_calculations").
			Field("employee_id").
			Unique().
			Required(),
	}
}

// Mixin for shared fields
func (SalaryCalculation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseFieldMixin{},
	}
}

// Indexes of the SalaryCalculation.
func (SalaryCalculation) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("employee_id", "calculation_month").Unique(),
		index.Fields("calculation_month"),
		index.Fields("employee_id"),
	}
}

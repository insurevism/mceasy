// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mceasy/ent/employee"
	"mceasy/ent/salarycalculation"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SalaryCalculationCreate is the builder for creating a SalaryCalculation entity.
type SalaryCalculationCreate struct {
	config
	mutation *SalaryCalculationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (scc *SalaryCalculationCreate) SetCreatedAt(t time.Time) *SalaryCalculationCreate {
	scc.mutation.SetCreatedAt(t)
	return scc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (scc *SalaryCalculationCreate) SetNillableCreatedAt(t *time.Time) *SalaryCalculationCreate {
	if t != nil {
		scc.SetCreatedAt(*t)
	}
	return scc
}

// SetModifiedAt sets the "modified_at" field.
func (scc *SalaryCalculationCreate) SetModifiedAt(t time.Time) *SalaryCalculationCreate {
	scc.mutation.SetModifiedAt(t)
	return scc
}

// SetNillableModifiedAt sets the "modified_at" field if the given value is not nil.
func (scc *SalaryCalculationCreate) SetNillableModifiedAt(t *time.Time) *SalaryCalculationCreate {
	if t != nil {
		scc.SetModifiedAt(*t)
	}
	return scc
}

// SetDeletedAt sets the "deleted_at" field.
func (scc *SalaryCalculationCreate) SetDeletedAt(t time.Time) *SalaryCalculationCreate {
	scc.mutation.SetDeletedAt(t)
	return scc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (scc *SalaryCalculationCreate) SetNillableDeletedAt(t *time.Time) *SalaryCalculationCreate {
	if t != nil {
		scc.SetDeletedAt(*t)
	}
	return scc
}

// SetEmployeeID sets the "employee_id" field.
func (scc *SalaryCalculationCreate) SetEmployeeID(u uint64) *SalaryCalculationCreate {
	scc.mutation.SetEmployeeID(u)
	return scc
}

// SetCalculationMonth sets the "calculation_month" field.
func (scc *SalaryCalculationCreate) SetCalculationMonth(t time.Time) *SalaryCalculationCreate {
	scc.mutation.SetCalculationMonth(t)
	return scc
}

// SetBaseSalary sets the "base_salary" field.
func (scc *SalaryCalculationCreate) SetBaseSalary(f float64) *SalaryCalculationCreate {
	scc.mutation.SetBaseSalary(f)
	return scc
}

// SetTotalWorkingDays sets the "total_working_days" field.
func (scc *SalaryCalculationCreate) SetTotalWorkingDays(i int) *SalaryCalculationCreate {
	scc.mutation.SetTotalWorkingDays(i)
	return scc
}

// SetAbsentDays sets the "absent_days" field.
func (scc *SalaryCalculationCreate) SetAbsentDays(i int) *SalaryCalculationCreate {
	scc.mutation.SetAbsentDays(i)
	return scc
}

// SetNillableAbsentDays sets the "absent_days" field if the given value is not nil.
func (scc *SalaryCalculationCreate) SetNillableAbsentDays(i *int) *SalaryCalculationCreate {
	if i != nil {
		scc.SetAbsentDays(*i)
	}
	return scc
}

// SetPresentDays sets the "present_days" field.
func (scc *SalaryCalculationCreate) SetPresentDays(i int) *SalaryCalculationCreate {
	scc.mutation.SetPresentDays(i)
	return scc
}

// SetNillablePresentDays sets the "present_days" field if the given value is not nil.
func (scc *SalaryCalculationCreate) SetNillablePresentDays(i *int) *SalaryCalculationCreate {
	if i != nil {
		scc.SetPresentDays(*i)
	}
	return scc
}

// SetFinalSalary sets the "final_salary" field.
func (scc *SalaryCalculationCreate) SetFinalSalary(f float64) *SalaryCalculationCreate {
	scc.mutation.SetFinalSalary(f)
	return scc
}

// SetDeductionAmount sets the "deduction_amount" field.
func (scc *SalaryCalculationCreate) SetDeductionAmount(f float64) *SalaryCalculationCreate {
	scc.mutation.SetDeductionAmount(f)
	return scc
}

// SetNillableDeductionAmount sets the "deduction_amount" field if the given value is not nil.
func (scc *SalaryCalculationCreate) SetNillableDeductionAmount(f *float64) *SalaryCalculationCreate {
	if f != nil {
		scc.SetDeductionAmount(*f)
	}
	return scc
}

// SetCalculationFormula sets the "calculation_formula" field.
func (scc *SalaryCalculationCreate) SetCalculationFormula(s string) *SalaryCalculationCreate {
	scc.mutation.SetCalculationFormula(s)
	return scc
}

// SetNillableCalculationFormula sets the "calculation_formula" field if the given value is not nil.
func (scc *SalaryCalculationCreate) SetNillableCalculationFormula(s *string) *SalaryCalculationCreate {
	if s != nil {
		scc.SetCalculationFormula(*s)
	}
	return scc
}

// SetID sets the "id" field.
func (scc *SalaryCalculationCreate) SetID(u uint64) *SalaryCalculationCreate {
	scc.mutation.SetID(u)
	return scc
}

// SetEmployee sets the "employee" edge to the Employee entity.
func (scc *SalaryCalculationCreate) SetEmployee(e *Employee) *SalaryCalculationCreate {
	return scc.SetEmployeeID(e.ID)
}

// Mutation returns the SalaryCalculationMutation object of the builder.
func (scc *SalaryCalculationCreate) Mutation() *SalaryCalculationMutation {
	return scc.mutation
}

// Save creates the SalaryCalculation in the database.
func (scc *SalaryCalculationCreate) Save(ctx context.Context) (*SalaryCalculation, error) {
	scc.defaults()
	return withHooks(ctx, scc.sqlSave, scc.mutation, scc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (scc *SalaryCalculationCreate) SaveX(ctx context.Context) *SalaryCalculation {
	v, err := scc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scc *SalaryCalculationCreate) Exec(ctx context.Context) error {
	_, err := scc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scc *SalaryCalculationCreate) ExecX(ctx context.Context) {
	if err := scc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scc *SalaryCalculationCreate) defaults() {
	if _, ok := scc.mutation.CreatedAt(); !ok {
		v := salarycalculation.DefaultCreatedAt()
		scc.mutation.SetCreatedAt(v)
	}
	if _, ok := scc.mutation.ModifiedAt(); !ok {
		v := salarycalculation.DefaultModifiedAt()
		scc.mutation.SetModifiedAt(v)
	}
	if _, ok := scc.mutation.AbsentDays(); !ok {
		v := salarycalculation.DefaultAbsentDays
		scc.mutation.SetAbsentDays(v)
	}
	if _, ok := scc.mutation.PresentDays(); !ok {
		v := salarycalculation.DefaultPresentDays
		scc.mutation.SetPresentDays(v)
	}
	if _, ok := scc.mutation.DeductionAmount(); !ok {
		v := salarycalculation.DefaultDeductionAmount
		scc.mutation.SetDeductionAmount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scc *SalaryCalculationCreate) check() error {
	if _, ok := scc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SalaryCalculation.created_at"`)}
	}
	if _, ok := scc.mutation.ModifiedAt(); !ok {
		return &ValidationError{Name: "modified_at", err: errors.New(`ent: missing required field "SalaryCalculation.modified_at"`)}
	}
	if _, ok := scc.mutation.EmployeeID(); !ok {
		return &ValidationError{Name: "employee_id", err: errors.New(`ent: missing required field "SalaryCalculation.employee_id"`)}
	}
	if _, ok := scc.mutation.CalculationMonth(); !ok {
		return &ValidationError{Name: "calculation_month", err: errors.New(`ent: missing required field "SalaryCalculation.calculation_month"`)}
	}
	if _, ok := scc.mutation.BaseSalary(); !ok {
		return &ValidationError{Name: "base_salary", err: errors.New(`ent: missing required field "SalaryCalculation.base_salary"`)}
	}
	if _, ok := scc.mutation.TotalWorkingDays(); !ok {
		return &ValidationError{Name: "total_working_days", err: errors.New(`ent: missing required field "SalaryCalculation.total_working_days"`)}
	}
	if _, ok := scc.mutation.AbsentDays(); !ok {
		return &ValidationError{Name: "absent_days", err: errors.New(`ent: missing required field "SalaryCalculation.absent_days"`)}
	}
	if _, ok := scc.mutation.PresentDays(); !ok {
		return &ValidationError{Name: "present_days", err: errors.New(`ent: missing required field "SalaryCalculation.present_days"`)}
	}
	if _, ok := scc.mutation.FinalSalary(); !ok {
		return &ValidationError{Name: "final_salary", err: errors.New(`ent: missing required field "SalaryCalculation.final_salary"`)}
	}
	if _, ok := scc.mutation.DeductionAmount(); !ok {
		return &ValidationError{Name: "deduction_amount", err: errors.New(`ent: missing required field "SalaryCalculation.deduction_amount"`)}
	}
	if _, ok := scc.mutation.EmployeeID(); !ok {
		return &ValidationError{Name: "employee", err: errors.New(`ent: missing required edge "SalaryCalculation.employee"`)}
	}
	return nil
}

func (scc *SalaryCalculationCreate) sqlSave(ctx context.Context) (*SalaryCalculation, error) {
	if err := scc.check(); err != nil {
		return nil, err
	}
	_node, _spec := scc.createSpec()
	if err := sqlgraph.CreateNode(ctx, scc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	scc.mutation.id = &_node.ID
	scc.mutation.done = true
	return _node, nil
}

func (scc *SalaryCalculationCreate) createSpec() (*SalaryCalculation, *sqlgraph.CreateSpec) {
	var (
		_node = &SalaryCalculation{config: scc.config}
		_spec = sqlgraph.NewCreateSpec(salarycalculation.Table, sqlgraph.NewFieldSpec(salarycalculation.FieldID, field.TypeUint64))
	)
	if id, ok := scc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := scc.mutation.CreatedAt(); ok {
		_spec.SetField(salarycalculation.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := scc.mutation.ModifiedAt(); ok {
		_spec.SetField(salarycalculation.FieldModifiedAt, field.TypeTime, value)
		_node.ModifiedAt = value
	}
	if value, ok := scc.mutation.DeletedAt(); ok {
		_spec.SetField(salarycalculation.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := scc.mutation.CalculationMonth(); ok {
		_spec.SetField(salarycalculation.FieldCalculationMonth, field.TypeTime, value)
		_node.CalculationMonth = value
	}
	if value, ok := scc.mutation.BaseSalary(); ok {
		_spec.SetField(salarycalculation.FieldBaseSalary, field.TypeFloat64, value)
		_node.BaseSalary = value
	}
	if value, ok := scc.mutation.TotalWorkingDays(); ok {
		_spec.SetField(salarycalculation.FieldTotalWorkingDays, field.TypeInt, value)
		_node.TotalWorkingDays = value
	}
	if value, ok := scc.mutation.AbsentDays(); ok {
		_spec.SetField(salarycalculation.FieldAbsentDays, field.TypeInt, value)
		_node.AbsentDays = value
	}
	if value, ok := scc.mutation.PresentDays(); ok {
		_spec.SetField(salarycalculation.FieldPresentDays, field.TypeInt, value)
		_node.PresentDays = value
	}
	if value, ok := scc.mutation.FinalSalary(); ok {
		_spec.SetField(salarycalculation.FieldFinalSalary, field.TypeFloat64, value)
		_node.FinalSalary = value
	}
	if value, ok := scc.mutation.DeductionAmount(); ok {
		_spec.SetField(salarycalculation.FieldDeductionAmount, field.TypeFloat64, value)
		_node.DeductionAmount = value
	}
	if value, ok := scc.mutation.CalculationFormula(); ok {
		_spec.SetField(salarycalculation.FieldCalculationFormula, field.TypeString, value)
		_node.CalculationFormula = value
	}
	if nodes := scc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   salarycalculation.EmployeeTable,
			Columns: []string{salarycalculation.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(employee.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EmployeeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SalaryCalculationCreateBulk is the builder for creating many SalaryCalculation entities in bulk.
type SalaryCalculationCreateBulk struct {
	config
	builders []*SalaryCalculationCreate
}

// Save creates the SalaryCalculation entities in the database.
func (sccb *SalaryCalculationCreateBulk) Save(ctx context.Context) ([]*SalaryCalculation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sccb.builders))
	nodes := make([]*SalaryCalculation, len(sccb.builders))
	mutators := make([]Mutator, len(sccb.builders))
	for i := range sccb.builders {
		func(i int, root context.Context) {
			builder := sccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SalaryCalculationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, sccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, sccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sccb *SalaryCalculationCreateBulk) SaveX(ctx context.Context) []*SalaryCalculation {
	v, err := sccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sccb *SalaryCalculationCreateBulk) Exec(ctx context.Context) error {
	_, err := sccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sccb *SalaryCalculationCreateBulk) ExecX(ctx context.Context) {
	if err := sccb.Exec(ctx); err != nil {
		panic(err)
	}
}

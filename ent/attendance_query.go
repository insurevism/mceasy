// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"mceasy/ent/attendance"
	"mceasy/ent/employee"
	"mceasy/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AttendanceQuery is the builder for querying Attendance entities.
type AttendanceQuery struct {
	config
	ctx          *QueryContext
	order        []attendance.OrderOption
	inters       []Interceptor
	predicates   []predicate.Attendance
	withEmployee *EmployeeQuery
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttendanceQuery builder.
func (aq *AttendanceQuery) Where(ps ...predicate.Attendance) *AttendanceQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AttendanceQuery) Limit(limit int) *AttendanceQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AttendanceQuery) Offset(offset int) *AttendanceQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AttendanceQuery) Unique(unique bool) *AttendanceQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AttendanceQuery) Order(o ...attendance.OrderOption) *AttendanceQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryEmployee chains the current query on the "employee" edge.
func (aq *AttendanceQuery) QueryEmployee() *EmployeeQuery {
	query := (&EmployeeClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attendance.Table, attendance.FieldID, selector),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attendance.EmployeeTable, attendance.EmployeeColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Attendance entity from the query.
// Returns a *NotFoundError when no Attendance was found.
func (aq *AttendanceQuery) First(ctx context.Context) (*Attendance, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attendance.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AttendanceQuery) FirstX(ctx context.Context) *Attendance {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Attendance ID from the query.
// Returns a *NotFoundError when no Attendance ID was found.
func (aq *AttendanceQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attendance.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AttendanceQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Attendance entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Attendance entity is found.
// Returns a *NotFoundError when no Attendance entities are found.
func (aq *AttendanceQuery) Only(ctx context.Context) (*Attendance, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attendance.Label}
	default:
		return nil, &NotSingularError{attendance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AttendanceQuery) OnlyX(ctx context.Context) *Attendance {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Attendance ID in the query.
// Returns a *NotSingularError when more than one Attendance ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AttendanceQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attendance.Label}
	default:
		err = &NotSingularError{attendance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AttendanceQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Attendances.
func (aq *AttendanceQuery) All(ctx context.Context) ([]*Attendance, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Attendance, *AttendanceQuery]()
	return withInterceptors[[]*Attendance](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AttendanceQuery) AllX(ctx context.Context) []*Attendance {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Attendance IDs.
func (aq *AttendanceQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(attendance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AttendanceQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AttendanceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AttendanceQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AttendanceQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AttendanceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AttendanceQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttendanceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AttendanceQuery) Clone() *AttendanceQuery {
	if aq == nil {
		return nil
	}
	return &AttendanceQuery{
		config:       aq.config,
		ctx:          aq.ctx.Clone(),
		order:        append([]attendance.OrderOption{}, aq.order...),
		inters:       append([]Interceptor{}, aq.inters...),
		predicates:   append([]predicate.Attendance{}, aq.predicates...),
		withEmployee: aq.withEmployee.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithEmployee tells the query-builder to eager-load the nodes that are connected to
// the "employee" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttendanceQuery) WithEmployee(opts ...func(*EmployeeQuery)) *AttendanceQuery {
	query := (&EmployeeClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withEmployee = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Attendance.Query().
//		GroupBy(attendance.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AttendanceQuery) GroupBy(field string, fields ...string) *AttendanceGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AttendanceGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = attendance.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Attendance.Query().
//		Select(attendance.FieldCreatedAt).
//		Scan(ctx, &v)
func (aq *AttendanceQuery) Select(fields ...string) *AttendanceSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AttendanceSelect{AttendanceQuery: aq}
	sbuild.label = attendance.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AttendanceSelect configured with the given aggregations.
func (aq *AttendanceQuery) Aggregate(fns ...AggregateFunc) *AttendanceSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AttendanceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !attendance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AttendanceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Attendance, error) {
	var (
		nodes       = []*Attendance{}
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withEmployee != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Attendance).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Attendance{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withEmployee; query != nil {
		if err := aq.loadEmployee(ctx, query, nodes, nil,
			func(n *Attendance, e *Employee) { n.Edges.Employee = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AttendanceQuery) loadEmployee(ctx context.Context, query *EmployeeQuery, nodes []*Attendance, init func(*Attendance), assign func(*Attendance, *Employee)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*Attendance)
	for i := range nodes {
		fk := nodes[i].EmployeeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(employee.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "employee_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aq *AttendanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AttendanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(attendance.Table, attendance.Columns, sqlgraph.NewFieldSpec(attendance.FieldID, field.TypeUint64))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attendance.FieldID)
		for i := range fields {
			if fields[i] != attendance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if aq.withEmployee != nil {
			_spec.Node.AddColumnOnce(attendance.FieldEmployeeID)
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AttendanceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(attendance.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = attendance.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range aq.modifiers {
		m(selector)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aq *AttendanceQuery) Modify(modifiers ...func(s *sql.Selector)) *AttendanceSelect {
	aq.modifiers = append(aq.modifiers, modifiers...)
	return aq.Select()
}

// AttendanceGroupBy is the group-by builder for Attendance entities.
type AttendanceGroupBy struct {
	selector
	build *AttendanceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AttendanceGroupBy) Aggregate(fns ...AggregateFunc) *AttendanceGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AttendanceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttendanceQuery, *AttendanceGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AttendanceGroupBy) sqlScan(ctx context.Context, root *AttendanceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AttendanceSelect is the builder for selecting fields of Attendance entities.
type AttendanceSelect struct {
	*AttendanceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AttendanceSelect) Aggregate(fns ...AggregateFunc) *AttendanceSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AttendanceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttendanceQuery, *AttendanceSelect](ctx, as.AttendanceQuery, as, as.inters, v)
}

func (as *AttendanceSelect) sqlScan(ctx context.Context, root *AttendanceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (as *AttendanceSelect) Modify(modifiers ...func(s *sql.Selector)) *AttendanceSelect {
	as.modifiers = append(as.modifiers, modifiers...)
	return as
}

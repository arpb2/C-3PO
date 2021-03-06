// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/arpb2/C-3PO/third_party/ent/level"
	"github.com/arpb2/C-3PO/third_party/ent/predicate"
	"github.com/arpb2/C-3PO/third_party/ent/user"
	"github.com/arpb2/C-3PO/third_party/ent/userlevel"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UserLevelQuery is the builder for querying UserLevel entities.
type UserLevelQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.UserLevel
	// eager-loading edges.
	withDeveloper *UserQuery
	withLevel     *LevelQuery
	withFKs       bool
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (ulq *UserLevelQuery) Where(ps ...predicate.UserLevel) *UserLevelQuery {
	ulq.predicates = append(ulq.predicates, ps...)
	return ulq
}

// Limit adds a limit step to the query.
func (ulq *UserLevelQuery) Limit(limit int) *UserLevelQuery {
	ulq.limit = &limit
	return ulq
}

// Offset adds an offset step to the query.
func (ulq *UserLevelQuery) Offset(offset int) *UserLevelQuery {
	ulq.offset = &offset
	return ulq
}

// Order adds an order step to the query.
func (ulq *UserLevelQuery) Order(o ...Order) *UserLevelQuery {
	ulq.order = append(ulq.order, o...)
	return ulq
}

// QueryDeveloper chains the current query on the developer edge.
func (ulq *UserLevelQuery) QueryDeveloper() *UserQuery {
	query := &UserQuery{config: ulq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(userlevel.Table, userlevel.FieldID, ulq.sqlQuery()),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, userlevel.DeveloperTable, userlevel.DeveloperColumn),
	)
	query.sql = sqlgraph.SetNeighbors(ulq.driver.Dialect(), step)
	return query
}

// QueryLevel chains the current query on the level edge.
func (ulq *UserLevelQuery) QueryLevel() *LevelQuery {
	query := &LevelQuery{config: ulq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(userlevel.Table, userlevel.FieldID, ulq.sqlQuery()),
		sqlgraph.To(level.Table, level.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, userlevel.LevelTable, userlevel.LevelColumn),
	)
	query.sql = sqlgraph.SetNeighbors(ulq.driver.Dialect(), step)
	return query
}

// First returns the first UserLevel entity in the query. Returns *NotFoundError when no userlevel was found.
func (ulq *UserLevelQuery) First(ctx context.Context) (*UserLevel, error) {
	uls, err := ulq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(uls) == 0 {
		return nil, &NotFoundError{userlevel.Label}
	}
	return uls[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ulq *UserLevelQuery) FirstX(ctx context.Context) *UserLevel {
	ul, err := ulq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ul
}

// FirstID returns the first UserLevel id in the query. Returns *NotFoundError when no id was found.
func (ulq *UserLevelQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ulq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userlevel.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (ulq *UserLevelQuery) FirstXID(ctx context.Context) int {
	id, err := ulq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only UserLevel entity in the query, returns an error if not exactly one entity was returned.
func (ulq *UserLevelQuery) Only(ctx context.Context) (*UserLevel, error) {
	uls, err := ulq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(uls) {
	case 1:
		return uls[0], nil
	case 0:
		return nil, &NotFoundError{userlevel.Label}
	default:
		return nil, &NotSingularError{userlevel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ulq *UserLevelQuery) OnlyX(ctx context.Context) *UserLevel {
	ul, err := ulq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ul
}

// OnlyID returns the only UserLevel id in the query, returns an error if not exactly one id was returned.
func (ulq *UserLevelQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ulq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userlevel.Label}
	default:
		err = &NotSingularError{userlevel.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (ulq *UserLevelQuery) OnlyXID(ctx context.Context) int {
	id, err := ulq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserLevels.
func (ulq *UserLevelQuery) All(ctx context.Context) ([]*UserLevel, error) {
	return ulq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ulq *UserLevelQuery) AllX(ctx context.Context) []*UserLevel {
	uls, err := ulq.All(ctx)
	if err != nil {
		panic(err)
	}
	return uls
}

// IDs executes the query and returns a list of UserLevel ids.
func (ulq *UserLevelQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ulq.Select(userlevel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ulq *UserLevelQuery) IDsX(ctx context.Context) []int {
	ids, err := ulq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ulq *UserLevelQuery) Count(ctx context.Context) (int, error) {
	return ulq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ulq *UserLevelQuery) CountX(ctx context.Context) int {
	count, err := ulq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ulq *UserLevelQuery) Exist(ctx context.Context) (bool, error) {
	return ulq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ulq *UserLevelQuery) ExistX(ctx context.Context) bool {
	exist, err := ulq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ulq *UserLevelQuery) Clone() *UserLevelQuery {
	return &UserLevelQuery{
		config:     ulq.config,
		limit:      ulq.limit,
		offset:     ulq.offset,
		order:      append([]Order{}, ulq.order...),
		unique:     append([]string{}, ulq.unique...),
		predicates: append([]predicate.UserLevel{}, ulq.predicates...),
		// clone intermediate query.
		sql: ulq.sql.Clone(),
	}
}

//  WithDeveloper tells the query-builder to eager-loads the nodes that are connected to
// the "developer" edge. The optional arguments used to configure the query builder of the edge.
func (ulq *UserLevelQuery) WithDeveloper(opts ...func(*UserQuery)) *UserLevelQuery {
	query := &UserQuery{config: ulq.config}
	for _, opt := range opts {
		opt(query)
	}
	ulq.withDeveloper = query
	return ulq
}

//  WithLevel tells the query-builder to eager-loads the nodes that are connected to
// the "level" edge. The optional arguments used to configure the query builder of the edge.
func (ulq *UserLevelQuery) WithLevel(opts ...func(*LevelQuery)) *UserLevelQuery {
	query := &LevelQuery{config: ulq.config}
	for _, opt := range opts {
		opt(query)
	}
	ulq.withLevel = query
	return ulq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserLevel.Query().
//		GroupBy(userlevel.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ulq *UserLevelQuery) GroupBy(field string, fields ...string) *UserLevelGroupBy {
	group := &UserLevelGroupBy{config: ulq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = ulq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.UserLevel.Query().
//		Select(userlevel.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (ulq *UserLevelQuery) Select(field string, fields ...string) *UserLevelSelect {
	selector := &UserLevelSelect{config: ulq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = ulq.sqlQuery()
	return selector
}

func (ulq *UserLevelQuery) sqlAll(ctx context.Context) ([]*UserLevel, error) {
	var (
		nodes       = []*UserLevel{}
		withFKs     = ulq.withFKs
		_spec       = ulq.querySpec()
		loadedTypes = [2]bool{
			ulq.withDeveloper != nil,
			ulq.withLevel != nil,
		}
	)
	if ulq.withDeveloper != nil || ulq.withLevel != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, userlevel.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &UserLevel{config: ulq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, ulq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ulq.withDeveloper; query != nil {
		ids := make([]uint, 0, len(nodes))
		nodeids := make(map[uint][]*UserLevel)
		for i := range nodes {
			if fk := nodes[i].user_level_developer; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_level_developer" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Developer = n
			}
		}
	}

	if query := ulq.withLevel; query != nil {
		ids := make([]uint, 0, len(nodes))
		nodeids := make(map[uint][]*UserLevel)
		for i := range nodes {
			if fk := nodes[i].user_level_level; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(level.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_level_level" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Level = n
			}
		}
	}

	return nodes, nil
}

func (ulq *UserLevelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ulq.querySpec()
	return sqlgraph.CountNodes(ctx, ulq.driver, _spec)
}

func (ulq *UserLevelQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ulq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ulq *UserLevelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userlevel.Table,
			Columns: userlevel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userlevel.FieldID,
			},
		},
		From:   ulq.sql,
		Unique: true,
	}
	if ps := ulq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ulq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ulq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ulq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ulq *UserLevelQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ulq.driver.Dialect())
	t1 := builder.Table(userlevel.Table)
	selector := builder.Select(t1.Columns(userlevel.Columns...)...).From(t1)
	if ulq.sql != nil {
		selector = ulq.sql
		selector.Select(selector.Columns(userlevel.Columns...)...)
	}
	for _, p := range ulq.predicates {
		p(selector)
	}
	for _, p := range ulq.order {
		p(selector)
	}
	if offset := ulq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ulq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserLevelGroupBy is the builder for group-by UserLevel entities.
type UserLevelGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ulgb *UserLevelGroupBy) Aggregate(fns ...Aggregate) *UserLevelGroupBy {
	ulgb.fns = append(ulgb.fns, fns...)
	return ulgb
}

// Scan applies the group-by query and scan the result into the given value.
func (ulgb *UserLevelGroupBy) Scan(ctx context.Context, v interface{}) error {
	return ulgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ulgb *UserLevelGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ulgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ulgb *UserLevelGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ulgb.fields) > 1 {
		return nil, errors.New("ent: UserLevelGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ulgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ulgb *UserLevelGroupBy) StringsX(ctx context.Context) []string {
	v, err := ulgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ulgb *UserLevelGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ulgb.fields) > 1 {
		return nil, errors.New("ent: UserLevelGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ulgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ulgb *UserLevelGroupBy) IntsX(ctx context.Context) []int {
	v, err := ulgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ulgb *UserLevelGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ulgb.fields) > 1 {
		return nil, errors.New("ent: UserLevelGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ulgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ulgb *UserLevelGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ulgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ulgb *UserLevelGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ulgb.fields) > 1 {
		return nil, errors.New("ent: UserLevelGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ulgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ulgb *UserLevelGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ulgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ulgb *UserLevelGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ulgb.sqlQuery().Query()
	if err := ulgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ulgb *UserLevelGroupBy) sqlQuery() *sql.Selector {
	selector := ulgb.sql
	columns := make([]string, 0, len(ulgb.fields)+len(ulgb.fns))
	columns = append(columns, ulgb.fields...)
	for _, fn := range ulgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(ulgb.fields...)
}

// UserLevelSelect is the builder for select fields of UserLevel entities.
type UserLevelSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (uls *UserLevelSelect) Scan(ctx context.Context, v interface{}) error {
	return uls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uls *UserLevelSelect) ScanX(ctx context.Context, v interface{}) {
	if err := uls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (uls *UserLevelSelect) Strings(ctx context.Context) ([]string, error) {
	if len(uls.fields) > 1 {
		return nil, errors.New("ent: UserLevelSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uls *UserLevelSelect) StringsX(ctx context.Context) []string {
	v, err := uls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (uls *UserLevelSelect) Ints(ctx context.Context) ([]int, error) {
	if len(uls.fields) > 1 {
		return nil, errors.New("ent: UserLevelSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uls *UserLevelSelect) IntsX(ctx context.Context) []int {
	v, err := uls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (uls *UserLevelSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uls.fields) > 1 {
		return nil, errors.New("ent: UserLevelSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uls *UserLevelSelect) Float64sX(ctx context.Context) []float64 {
	v, err := uls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (uls *UserLevelSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uls.fields) > 1 {
		return nil, errors.New("ent: UserLevelSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uls *UserLevelSelect) BoolsX(ctx context.Context) []bool {
	v, err := uls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uls *UserLevelSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uls.sqlQuery().Query()
	if err := uls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uls *UserLevelSelect) sqlQuery() sql.Querier {
	selector := uls.sql
	selector.Select(selector.Columns(uls.fields...)...)
	return selector
}

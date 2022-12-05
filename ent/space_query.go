// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/dilliver/ent/folder"
	"github.com/ugent-library/dilliver/ent/predicate"
	"github.com/ugent-library/dilliver/ent/space"
)

// SpaceQuery is the builder for querying Space entities.
type SpaceQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	predicates  []predicate.Space
	withFolders *FolderQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SpaceQuery builder.
func (sq *SpaceQuery) Where(ps ...predicate.Space) *SpaceQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *SpaceQuery) Limit(limit int) *SpaceQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *SpaceQuery) Offset(offset int) *SpaceQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SpaceQuery) Unique(unique bool) *SpaceQuery {
	sq.unique = &unique
	return sq
}

// Order adds an order step to the query.
func (sq *SpaceQuery) Order(o ...OrderFunc) *SpaceQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryFolders chains the current query on the "folders" edge.
func (sq *SpaceQuery) QueryFolders() *FolderQuery {
	query := &FolderQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(space.Table, space.FieldID, selector),
			sqlgraph.To(folder.Table, folder.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, space.FoldersTable, space.FoldersColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Space entity from the query.
// Returns a *NotFoundError when no Space was found.
func (sq *SpaceQuery) First(ctx context.Context) (*Space, error) {
	nodes, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{space.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SpaceQuery) FirstX(ctx context.Context) *Space {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Space ID from the query.
// Returns a *NotFoundError when no Space ID was found.
func (sq *SpaceQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{space.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SpaceQuery) FirstIDX(ctx context.Context) string {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Space entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Space entity is found.
// Returns a *NotFoundError when no Space entities are found.
func (sq *SpaceQuery) Only(ctx context.Context) (*Space, error) {
	nodes, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{space.Label}
	default:
		return nil, &NotSingularError{space.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SpaceQuery) OnlyX(ctx context.Context) *Space {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Space ID in the query.
// Returns a *NotSingularError when more than one Space ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SpaceQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{space.Label}
	default:
		err = &NotSingularError{space.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SpaceQuery) OnlyIDX(ctx context.Context) string {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Spaces.
func (sq *SpaceQuery) All(ctx context.Context) ([]*Space, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *SpaceQuery) AllX(ctx context.Context) []*Space {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Space IDs.
func (sq *SpaceQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := sq.Select(space.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SpaceQuery) IDsX(ctx context.Context) []string {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SpaceQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SpaceQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SpaceQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SpaceQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SpaceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SpaceQuery) Clone() *SpaceQuery {
	if sq == nil {
		return nil
	}
	return &SpaceQuery{
		config:      sq.config,
		limit:       sq.limit,
		offset:      sq.offset,
		order:       append([]OrderFunc{}, sq.order...),
		predicates:  append([]predicate.Space{}, sq.predicates...),
		withFolders: sq.withFolders.Clone(),
		// clone intermediate query.
		sql:    sq.sql.Clone(),
		path:   sq.path,
		unique: sq.unique,
	}
}

// WithFolders tells the query-builder to eager-load the nodes that are connected to
// the "folders" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SpaceQuery) WithFolders(opts ...func(*FolderQuery)) *SpaceQuery {
	query := &FolderQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withFolders = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Space.Query().
//		GroupBy(space.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SpaceQuery) GroupBy(field string, fields ...string) *SpaceGroupBy {
	grbuild := &SpaceGroupBy{config: sq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(ctx), nil
	}
	grbuild.label = space.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Space.Query().
//		Select(space.FieldName).
//		Scan(ctx, &v)
func (sq *SpaceQuery) Select(fields ...string) *SpaceSelect {
	sq.fields = append(sq.fields, fields...)
	selbuild := &SpaceSelect{SpaceQuery: sq}
	selbuild.label = space.Label
	selbuild.flds, selbuild.scan = &sq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a SpaceSelect configured with the given aggregations.
func (sq *SpaceQuery) Aggregate(fns ...AggregateFunc) *SpaceSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SpaceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sq.fields {
		if !space.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SpaceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Space, error) {
	var (
		nodes       = []*Space{}
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withFolders != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Space).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Space{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withFolders; query != nil {
		if err := sq.loadFolders(ctx, query, nodes,
			func(n *Space) { n.Edges.Folders = []*Folder{} },
			func(n *Space, e *Folder) { n.Edges.Folders = append(n.Edges.Folders, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SpaceQuery) loadFolders(ctx context.Context, query *FolderQuery, nodes []*Space, init func(*Space), assign func(*Space, *Folder)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Space)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Folder(func(s *sql.Selector) {
		s.Where(sql.InValues(space.FoldersColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.space_folders
		if fk == nil {
			return fmt.Errorf(`foreign-key "space_folders" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "space_folders" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (sq *SpaceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.fields
	if len(sq.fields) > 0 {
		_spec.Unique = sq.unique != nil && *sq.unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SpaceQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (sq *SpaceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   space.Table,
			Columns: space.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: space.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, space.FieldID)
		for i := range fields {
			if fields[i] != space.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SpaceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(space.Table)
	columns := sq.fields
	if len(columns) == 0 {
		columns = space.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.unique != nil && *sq.unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SpaceGroupBy is the group-by builder for Space entities.
type SpaceGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SpaceGroupBy) Aggregate(fns ...AggregateFunc) *SpaceGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgb *SpaceGroupBy) Scan(ctx context.Context, v any) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

func (sgb *SpaceGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range sgb.fields {
		if !space.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *SpaceGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql.Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
		for _, f := range sgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgb.fields...)...)
}

// SpaceSelect is the builder for selecting fields of Space entities.
type SpaceSelect struct {
	*SpaceQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SpaceSelect) Aggregate(fns ...AggregateFunc) *SpaceSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SpaceSelect) Scan(ctx context.Context, v any) error {
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	ss.sql = ss.SpaceQuery.sqlQuery(ctx)
	return ss.sqlScan(ctx, v)
}

func (ss *SpaceSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(ss.sql))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ss.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ss.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ss.sql.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

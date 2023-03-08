// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/deliver/ent/folder"
	"github.com/ugent-library/deliver/ent/predicate"
	"github.com/ugent-library/deliver/ent/space"
)

// SpaceUpdate is the builder for updating Space entities.
type SpaceUpdate struct {
	config
	hooks    []Hook
	mutation *SpaceMutation
}

// Where appends a list predicates to the SpaceUpdate builder.
func (su *SpaceUpdate) Where(ps ...predicate.Space) *SpaceUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *SpaceUpdate) SetName(s string) *SpaceUpdate {
	su.mutation.SetName(s)
	return su
}

// SetAdmins sets the "admins" field.
func (su *SpaceUpdate) SetAdmins(s []string) *SpaceUpdate {
	su.mutation.SetAdmins(s)
	return su
}

// AppendAdmins appends s to the "admins" field.
func (su *SpaceUpdate) AppendAdmins(s []string) *SpaceUpdate {
	su.mutation.AppendAdmins(s)
	return su
}

// ClearAdmins clears the value of the "admins" field.
func (su *SpaceUpdate) ClearAdmins() *SpaceUpdate {
	su.mutation.ClearAdmins()
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SpaceUpdate) SetUpdatedAt(t time.Time) *SpaceUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// AddFolderIDs adds the "folders" edge to the Folder entity by IDs.
func (su *SpaceUpdate) AddFolderIDs(ids ...string) *SpaceUpdate {
	su.mutation.AddFolderIDs(ids...)
	return su
}

// AddFolders adds the "folders" edges to the Folder entity.
func (su *SpaceUpdate) AddFolders(f ...*Folder) *SpaceUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return su.AddFolderIDs(ids...)
}

// Mutation returns the SpaceMutation object of the builder.
func (su *SpaceUpdate) Mutation() *SpaceMutation {
	return su.mutation
}

// ClearFolders clears all "folders" edges to the Folder entity.
func (su *SpaceUpdate) ClearFolders() *SpaceUpdate {
	su.mutation.ClearFolders()
	return su
}

// RemoveFolderIDs removes the "folders" edge to Folder entities by IDs.
func (su *SpaceUpdate) RemoveFolderIDs(ids ...string) *SpaceUpdate {
	su.mutation.RemoveFolderIDs(ids...)
	return su
}

// RemoveFolders removes "folders" edges to Folder entities.
func (su *SpaceUpdate) RemoveFolders(f ...*Folder) *SpaceUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return su.RemoveFolderIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SpaceUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks[int, SpaceMutation](ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SpaceUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SpaceUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SpaceUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SpaceUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := space.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

func (su *SpaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(space.Table, space.Columns, sqlgraph.NewFieldSpec(space.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(space.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Admins(); ok {
		_spec.SetField(space.FieldAdmins, field.TypeJSON, value)
	}
	if value, ok := su.mutation.AppendedAdmins(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, space.FieldAdmins, value)
		})
	}
	if su.mutation.AdminsCleared() {
		_spec.ClearField(space.FieldAdmins, field.TypeJSON)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(space.FieldUpdatedAt, field.TypeTime, value)
	}
	if su.mutation.FoldersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   space.FoldersTable,
			Columns: []string{space.FoldersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: folder.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedFoldersIDs(); len(nodes) > 0 && !su.mutation.FoldersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   space.FoldersTable,
			Columns: []string{space.FoldersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: folder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.FoldersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   space.FoldersTable,
			Columns: []string{space.FoldersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: folder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{space.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SpaceUpdateOne is the builder for updating a single Space entity.
type SpaceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SpaceMutation
}

// SetName sets the "name" field.
func (suo *SpaceUpdateOne) SetName(s string) *SpaceUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetAdmins sets the "admins" field.
func (suo *SpaceUpdateOne) SetAdmins(s []string) *SpaceUpdateOne {
	suo.mutation.SetAdmins(s)
	return suo
}

// AppendAdmins appends s to the "admins" field.
func (suo *SpaceUpdateOne) AppendAdmins(s []string) *SpaceUpdateOne {
	suo.mutation.AppendAdmins(s)
	return suo
}

// ClearAdmins clears the value of the "admins" field.
func (suo *SpaceUpdateOne) ClearAdmins() *SpaceUpdateOne {
	suo.mutation.ClearAdmins()
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SpaceUpdateOne) SetUpdatedAt(t time.Time) *SpaceUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// AddFolderIDs adds the "folders" edge to the Folder entity by IDs.
func (suo *SpaceUpdateOne) AddFolderIDs(ids ...string) *SpaceUpdateOne {
	suo.mutation.AddFolderIDs(ids...)
	return suo
}

// AddFolders adds the "folders" edges to the Folder entity.
func (suo *SpaceUpdateOne) AddFolders(f ...*Folder) *SpaceUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return suo.AddFolderIDs(ids...)
}

// Mutation returns the SpaceMutation object of the builder.
func (suo *SpaceUpdateOne) Mutation() *SpaceMutation {
	return suo.mutation
}

// ClearFolders clears all "folders" edges to the Folder entity.
func (suo *SpaceUpdateOne) ClearFolders() *SpaceUpdateOne {
	suo.mutation.ClearFolders()
	return suo
}

// RemoveFolderIDs removes the "folders" edge to Folder entities by IDs.
func (suo *SpaceUpdateOne) RemoveFolderIDs(ids ...string) *SpaceUpdateOne {
	suo.mutation.RemoveFolderIDs(ids...)
	return suo
}

// RemoveFolders removes "folders" edges to Folder entities.
func (suo *SpaceUpdateOne) RemoveFolders(f ...*Folder) *SpaceUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return suo.RemoveFolderIDs(ids...)
}

// Where appends a list predicates to the SpaceUpdate builder.
func (suo *SpaceUpdateOne) Where(ps ...predicate.Space) *SpaceUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SpaceUpdateOne) Select(field string, fields ...string) *SpaceUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Space entity.
func (suo *SpaceUpdateOne) Save(ctx context.Context) (*Space, error) {
	suo.defaults()
	return withHooks[*Space, SpaceMutation](ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SpaceUpdateOne) SaveX(ctx context.Context) *Space {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SpaceUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SpaceUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SpaceUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := space.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

func (suo *SpaceUpdateOne) sqlSave(ctx context.Context) (_node *Space, err error) {
	_spec := sqlgraph.NewUpdateSpec(space.Table, space.Columns, sqlgraph.NewFieldSpec(space.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Space.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, space.FieldID)
		for _, f := range fields {
			if !space.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != space.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(space.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Admins(); ok {
		_spec.SetField(space.FieldAdmins, field.TypeJSON, value)
	}
	if value, ok := suo.mutation.AppendedAdmins(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, space.FieldAdmins, value)
		})
	}
	if suo.mutation.AdminsCleared() {
		_spec.ClearField(space.FieldAdmins, field.TypeJSON)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(space.FieldUpdatedAt, field.TypeTime, value)
	}
	if suo.mutation.FoldersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   space.FoldersTable,
			Columns: []string{space.FoldersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: folder.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedFoldersIDs(); len(nodes) > 0 && !suo.mutation.FoldersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   space.FoldersTable,
			Columns: []string{space.FoldersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: folder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.FoldersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   space.FoldersTable,
			Columns: []string{space.FoldersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: folder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Space{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{space.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/deliver/ent/file"
	"github.com/ugent-library/deliver/ent/folder"
	"github.com/ugent-library/deliver/ent/predicate"
	"github.com/ugent-library/deliver/ent/space"
)

// FolderUpdate is the builder for updating Folder entities.
type FolderUpdate struct {
	config
	hooks    []Hook
	mutation *FolderMutation
}

// Where appends a list predicates to the FolderUpdate builder.
func (fu *FolderUpdate) Where(ps ...predicate.Folder) *FolderUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetSpaceID sets the "space_id" field.
func (fu *FolderUpdate) SetSpaceID(s string) *FolderUpdate {
	fu.mutation.SetSpaceID(s)
	return fu
}

// SetName sets the "name" field.
func (fu *FolderUpdate) SetName(s string) *FolderUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FolderUpdate) SetUpdatedAt(t time.Time) *FolderUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetExpiresAt sets the "expires_at" field.
func (fu *FolderUpdate) SetExpiresAt(t time.Time) *FolderUpdate {
	fu.mutation.SetExpiresAt(t)
	return fu
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (fu *FolderUpdate) SetNillableExpiresAt(t *time.Time) *FolderUpdate {
	if t != nil {
		fu.SetExpiresAt(*t)
	}
	return fu
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (fu *FolderUpdate) ClearExpiresAt() *FolderUpdate {
	fu.mutation.ClearExpiresAt()
	return fu
}

// SetSpace sets the "space" edge to the Space entity.
func (fu *FolderUpdate) SetSpace(s *Space) *FolderUpdate {
	return fu.SetSpaceID(s.ID)
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (fu *FolderUpdate) AddFileIDs(ids ...string) *FolderUpdate {
	fu.mutation.AddFileIDs(ids...)
	return fu
}

// AddFiles adds the "files" edges to the File entity.
func (fu *FolderUpdate) AddFiles(f ...*File) *FolderUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.AddFileIDs(ids...)
}

// Mutation returns the FolderMutation object of the builder.
func (fu *FolderUpdate) Mutation() *FolderMutation {
	return fu.mutation
}

// ClearSpace clears the "space" edge to the Space entity.
func (fu *FolderUpdate) ClearSpace() *FolderUpdate {
	fu.mutation.ClearSpace()
	return fu
}

// ClearFiles clears all "files" edges to the File entity.
func (fu *FolderUpdate) ClearFiles() *FolderUpdate {
	fu.mutation.ClearFiles()
	return fu
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (fu *FolderUpdate) RemoveFileIDs(ids ...string) *FolderUpdate {
	fu.mutation.RemoveFileIDs(ids...)
	return fu
}

// RemoveFiles removes "files" edges to File entities.
func (fu *FolderUpdate) RemoveFiles(f ...*File) *FolderUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.RemoveFileIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FolderUpdate) Save(ctx context.Context) (int, error) {
	fu.defaults()
	return withHooks[int, FolderMutation](ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FolderUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FolderUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FolderUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FolderUpdate) defaults() {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		v := folder.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FolderUpdate) check() error {
	if _, ok := fu.mutation.SpaceID(); fu.mutation.SpaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Folder.space"`)
	}
	return nil
}

func (fu *FolderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   folder.Table,
			Columns: folder.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: folder.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.SetField(folder.FieldName, field.TypeString, value)
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(folder.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fu.mutation.ExpiresAt(); ok {
		_spec.SetField(folder.FieldExpiresAt, field.TypeTime, value)
	}
	if fu.mutation.ExpiresAtCleared() {
		_spec.ClearField(folder.FieldExpiresAt, field.TypeTime)
	}
	if fu.mutation.SpaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   folder.SpaceTable,
			Columns: []string{folder.SpaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: space.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.SpaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   folder.SpaceTable,
			Columns: []string{folder.SpaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: space.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.FilesTable,
			Columns: []string{folder.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: file.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedFilesIDs(); len(nodes) > 0 && !fu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.FilesTable,
			Columns: []string{folder.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.FilesTable,
			Columns: []string{folder.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{folder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FolderUpdateOne is the builder for updating a single Folder entity.
type FolderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FolderMutation
}

// SetSpaceID sets the "space_id" field.
func (fuo *FolderUpdateOne) SetSpaceID(s string) *FolderUpdateOne {
	fuo.mutation.SetSpaceID(s)
	return fuo
}

// SetName sets the "name" field.
func (fuo *FolderUpdateOne) SetName(s string) *FolderUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FolderUpdateOne) SetUpdatedAt(t time.Time) *FolderUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetExpiresAt sets the "expires_at" field.
func (fuo *FolderUpdateOne) SetExpiresAt(t time.Time) *FolderUpdateOne {
	fuo.mutation.SetExpiresAt(t)
	return fuo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (fuo *FolderUpdateOne) SetNillableExpiresAt(t *time.Time) *FolderUpdateOne {
	if t != nil {
		fuo.SetExpiresAt(*t)
	}
	return fuo
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (fuo *FolderUpdateOne) ClearExpiresAt() *FolderUpdateOne {
	fuo.mutation.ClearExpiresAt()
	return fuo
}

// SetSpace sets the "space" edge to the Space entity.
func (fuo *FolderUpdateOne) SetSpace(s *Space) *FolderUpdateOne {
	return fuo.SetSpaceID(s.ID)
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (fuo *FolderUpdateOne) AddFileIDs(ids ...string) *FolderUpdateOne {
	fuo.mutation.AddFileIDs(ids...)
	return fuo
}

// AddFiles adds the "files" edges to the File entity.
func (fuo *FolderUpdateOne) AddFiles(f ...*File) *FolderUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.AddFileIDs(ids...)
}

// Mutation returns the FolderMutation object of the builder.
func (fuo *FolderUpdateOne) Mutation() *FolderMutation {
	return fuo.mutation
}

// ClearSpace clears the "space" edge to the Space entity.
func (fuo *FolderUpdateOne) ClearSpace() *FolderUpdateOne {
	fuo.mutation.ClearSpace()
	return fuo
}

// ClearFiles clears all "files" edges to the File entity.
func (fuo *FolderUpdateOne) ClearFiles() *FolderUpdateOne {
	fuo.mutation.ClearFiles()
	return fuo
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (fuo *FolderUpdateOne) RemoveFileIDs(ids ...string) *FolderUpdateOne {
	fuo.mutation.RemoveFileIDs(ids...)
	return fuo
}

// RemoveFiles removes "files" edges to File entities.
func (fuo *FolderUpdateOne) RemoveFiles(f ...*File) *FolderUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.RemoveFileIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FolderUpdateOne) Select(field string, fields ...string) *FolderUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Folder entity.
func (fuo *FolderUpdateOne) Save(ctx context.Context) (*Folder, error) {
	fuo.defaults()
	return withHooks[*Folder, FolderMutation](ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FolderUpdateOne) SaveX(ctx context.Context) *Folder {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FolderUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FolderUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FolderUpdateOne) defaults() {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		v := folder.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FolderUpdateOne) check() error {
	if _, ok := fuo.mutation.SpaceID(); fuo.mutation.SpaceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Folder.space"`)
	}
	return nil
}

func (fuo *FolderUpdateOne) sqlSave(ctx context.Context) (_node *Folder, err error) {
	if err := fuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   folder.Table,
			Columns: folder.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: folder.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Folder.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, folder.FieldID)
		for _, f := range fields {
			if !folder.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != folder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Name(); ok {
		_spec.SetField(folder.FieldName, field.TypeString, value)
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(folder.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fuo.mutation.ExpiresAt(); ok {
		_spec.SetField(folder.FieldExpiresAt, field.TypeTime, value)
	}
	if fuo.mutation.ExpiresAtCleared() {
		_spec.ClearField(folder.FieldExpiresAt, field.TypeTime)
	}
	if fuo.mutation.SpaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   folder.SpaceTable,
			Columns: []string{folder.SpaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: space.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.SpaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   folder.SpaceTable,
			Columns: []string{folder.SpaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: space.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.FilesTable,
			Columns: []string{folder.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: file.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedFilesIDs(); len(nodes) > 0 && !fuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.FilesTable,
			Columns: []string{folder.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   folder.FilesTable,
			Columns: []string{folder.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Folder{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{folder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}

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
	"github.com/ugent-library/dilliver/ent/file"
	"github.com/ugent-library/dilliver/ent/folder"
	"github.com/ugent-library/dilliver/ent/predicate"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	hooks    []Hook
	mutation *FileMutation
}

// Where appends a list predicates to the FileUpdate builder.
func (fu *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetFolderID sets the "folder_id" field.
func (fu *FileUpdate) SetFolderID(s string) *FileUpdate {
	fu.mutation.SetFolderID(s)
	return fu
}

// SetMd5 sets the "md5" field.
func (fu *FileUpdate) SetMd5(s string) *FileUpdate {
	fu.mutation.SetMd5(s)
	return fu
}

// SetName sets the "name" field.
func (fu *FileUpdate) SetName(s string) *FileUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetSize sets the "size" field.
func (fu *FileUpdate) SetSize(i int64) *FileUpdate {
	fu.mutation.ResetSize()
	fu.mutation.SetSize(i)
	return fu
}

// AddSize adds i to the "size" field.
func (fu *FileUpdate) AddSize(i int64) *FileUpdate {
	fu.mutation.AddSize(i)
	return fu
}

// SetContentType sets the "content_type" field.
func (fu *FileUpdate) SetContentType(s string) *FileUpdate {
	fu.mutation.SetContentType(s)
	return fu
}

// SetDownloads sets the "downloads" field.
func (fu *FileUpdate) SetDownloads(i int32) *FileUpdate {
	fu.mutation.ResetDownloads()
	fu.mutation.SetDownloads(i)
	return fu
}

// SetNillableDownloads sets the "downloads" field if the given value is not nil.
func (fu *FileUpdate) SetNillableDownloads(i *int32) *FileUpdate {
	if i != nil {
		fu.SetDownloads(*i)
	}
	return fu
}

// AddDownloads adds i to the "downloads" field.
func (fu *FileUpdate) AddDownloads(i int32) *FileUpdate {
	fu.mutation.AddDownloads(i)
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FileUpdate) SetUpdatedAt(t time.Time) *FileUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetFolder sets the "folder" edge to the Folder entity.
func (fu *FileUpdate) SetFolder(f *Folder) *FileUpdate {
	return fu.SetFolderID(f.ID)
}

// Mutation returns the FileMutation object of the builder.
func (fu *FileUpdate) Mutation() *FileMutation {
	return fu.mutation
}

// ClearFolder clears the "folder" edge to the Folder entity.
func (fu *FileUpdate) ClearFolder() *FileUpdate {
	fu.mutation.ClearFolder()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	fu.defaults()
	if len(fu.hooks) == 0 {
		if err = fu.check(); err != nil {
			return 0, err
		}
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fu.check(); err != nil {
				return 0, err
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			if fu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FileUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FileUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FileUpdate) defaults() {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		v := file.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FileUpdate) check() error {
	if _, ok := fu.mutation.FolderID(); fu.mutation.FolderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "File.folder"`)
	}
	return nil
}

func (fu *FileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   file.Table,
			Columns: file.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: file.FieldID,
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
	if value, ok := fu.mutation.Md5(); ok {
		_spec.SetField(file.FieldMd5, field.TypeString, value)
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.SetField(file.FieldName, field.TypeString, value)
	}
	if value, ok := fu.mutation.Size(); ok {
		_spec.SetField(file.FieldSize, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.AddedSize(); ok {
		_spec.AddField(file.FieldSize, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.ContentType(); ok {
		_spec.SetField(file.FieldContentType, field.TypeString, value)
	}
	if value, ok := fu.mutation.Downloads(); ok {
		_spec.SetField(file.FieldDownloads, field.TypeInt32, value)
	}
	if value, ok := fu.mutation.AddedDownloads(); ok {
		_spec.AddField(file.FieldDownloads, field.TypeInt32, value)
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(file.FieldUpdatedAt, field.TypeTime, value)
	}
	if fu.mutation.FolderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.FolderTable,
			Columns: []string{file.FolderColumn},
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
	if nodes := fu.mutation.FolderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.FolderTable,
			Columns: []string{file.FolderColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileMutation
}

// SetFolderID sets the "folder_id" field.
func (fuo *FileUpdateOne) SetFolderID(s string) *FileUpdateOne {
	fuo.mutation.SetFolderID(s)
	return fuo
}

// SetMd5 sets the "md5" field.
func (fuo *FileUpdateOne) SetMd5(s string) *FileUpdateOne {
	fuo.mutation.SetMd5(s)
	return fuo
}

// SetName sets the "name" field.
func (fuo *FileUpdateOne) SetName(s string) *FileUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetSize sets the "size" field.
func (fuo *FileUpdateOne) SetSize(i int64) *FileUpdateOne {
	fuo.mutation.ResetSize()
	fuo.mutation.SetSize(i)
	return fuo
}

// AddSize adds i to the "size" field.
func (fuo *FileUpdateOne) AddSize(i int64) *FileUpdateOne {
	fuo.mutation.AddSize(i)
	return fuo
}

// SetContentType sets the "content_type" field.
func (fuo *FileUpdateOne) SetContentType(s string) *FileUpdateOne {
	fuo.mutation.SetContentType(s)
	return fuo
}

// SetDownloads sets the "downloads" field.
func (fuo *FileUpdateOne) SetDownloads(i int32) *FileUpdateOne {
	fuo.mutation.ResetDownloads()
	fuo.mutation.SetDownloads(i)
	return fuo
}

// SetNillableDownloads sets the "downloads" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableDownloads(i *int32) *FileUpdateOne {
	if i != nil {
		fuo.SetDownloads(*i)
	}
	return fuo
}

// AddDownloads adds i to the "downloads" field.
func (fuo *FileUpdateOne) AddDownloads(i int32) *FileUpdateOne {
	fuo.mutation.AddDownloads(i)
	return fuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FileUpdateOne) SetUpdatedAt(t time.Time) *FileUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetFolder sets the "folder" edge to the Folder entity.
func (fuo *FileUpdateOne) SetFolder(f *Folder) *FileUpdateOne {
	return fuo.SetFolderID(f.ID)
}

// Mutation returns the FileMutation object of the builder.
func (fuo *FileUpdateOne) Mutation() *FileMutation {
	return fuo.mutation
}

// ClearFolder clears the "folder" edge to the Folder entity.
func (fuo *FileUpdateOne) ClearFolder() *FileUpdateOne {
	fuo.mutation.ClearFolder()
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FileUpdateOne) Select(field string, fields ...string) *FileUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated File entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	var (
		err  error
		node *File
	)
	fuo.defaults()
	if len(fuo.hooks) == 0 {
		if err = fuo.check(); err != nil {
			return nil, err
		}
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fuo.check(); err != nil {
				return nil, err
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			if fuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*File)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FileMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FileUpdateOne) SaveX(ctx context.Context) *File {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FileUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FileUpdateOne) defaults() {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		v := file.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FileUpdateOne) check() error {
	if _, ok := fuo.mutation.FolderID(); fuo.mutation.FolderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "File.folder"`)
	}
	return nil
}

func (fuo *FileUpdateOne) sqlSave(ctx context.Context) (_node *File, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   file.Table,
			Columns: file.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: file.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "File.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, file.FieldID)
		for _, f := range fields {
			if !file.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != file.FieldID {
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
	if value, ok := fuo.mutation.Md5(); ok {
		_spec.SetField(file.FieldMd5, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Name(); ok {
		_spec.SetField(file.FieldName, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Size(); ok {
		_spec.SetField(file.FieldSize, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.AddedSize(); ok {
		_spec.AddField(file.FieldSize, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.ContentType(); ok {
		_spec.SetField(file.FieldContentType, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Downloads(); ok {
		_spec.SetField(file.FieldDownloads, field.TypeInt32, value)
	}
	if value, ok := fuo.mutation.AddedDownloads(); ok {
		_spec.AddField(file.FieldDownloads, field.TypeInt32, value)
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(file.FieldUpdatedAt, field.TypeTime, value)
	}
	if fuo.mutation.FolderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.FolderTable,
			Columns: []string{file.FolderColumn},
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
	if nodes := fuo.mutation.FolderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.FolderTable,
			Columns: []string{file.FolderColumn},
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
	_node = &File{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

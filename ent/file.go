// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/ugent-library/dilliver/ent/file"
	"github.com/ugent-library/dilliver/ent/folder"
)

// File is the model entity for the File schema.
type File struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges        FileEdges `json:"edges"`
	folder_files *string
}

// FileEdges holds the relations/edges for other nodes in the graph.
type FileEdges struct {
	// Folder holds the value of the folder edge.
	Folder *Folder `json:"folder,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FolderOrErr returns the Folder value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) FolderOrErr() (*Folder, error) {
	if e.loadedTypes[0] {
		if e.Folder == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: folder.Label}
		}
		return e.Folder, nil
	}
	return nil, &NotLoadedError{edge: "folder"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*File) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case file.FieldID:
			values[i] = new(sql.NullString)
		case file.ForeignKeys[0]: // folder_files
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type File", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case file.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				f.ID = value.String
			}
		case file.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field folder_files", values[i])
			} else if value.Valid {
				f.folder_files = new(string)
				*f.folder_files = value.String
			}
		}
	}
	return nil
}

// QueryFolder queries the "folder" edge of the File entity.
func (f *File) QueryFolder() *FolderQuery {
	return (&FileClient{config: f.config}).QueryFolder(f)
}

// Update returns a builder for updating this File.
// Note that you need to call File.Unwrap() before calling this method if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return (&FileClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the File entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Files is a parsable slice of File.
type Files []*File

func (f Files) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}

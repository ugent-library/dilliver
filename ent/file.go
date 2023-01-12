// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ugent-library/deliver/ent/file"
	"github.com/ugent-library/deliver/ent/folder"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// FolderID holds the value of the "folder_id" field.
	FolderID string `json:"folder_id,omitempty"`
	// Md5 holds the value of the "md5" field.
	Md5 string `json:"md5,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Size holds the value of the "size" field.
	Size int64 `json:"size,omitempty"`
	// ContentType holds the value of the "content_type" field.
	ContentType string `json:"content_type,omitempty"`
	// Downloads holds the value of the "downloads" field.
	Downloads int64 `json:"downloads,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges FileEdges `json:"edges"`
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
		case file.FieldSize, file.FieldDownloads:
			values[i] = new(sql.NullInt64)
		case file.FieldID, file.FieldFolderID, file.FieldMd5, file.FieldName, file.FieldContentType:
			values[i] = new(sql.NullString)
		case file.FieldCreatedAt, file.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
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
		case file.FieldFolderID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field folder_id", values[i])
			} else if value.Valid {
				f.FolderID = value.String
			}
		case file.FieldMd5:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field md5", values[i])
			} else if value.Valid {
				f.Md5 = value.String
			}
		case file.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				f.Name = value.String
			}
		case file.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				f.Size = value.Int64
			}
		case file.FieldContentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content_type", values[i])
			} else if value.Valid {
				f.ContentType = value.String
			}
		case file.FieldDownloads:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field downloads", values[i])
			} else if value.Valid {
				f.Downloads = value.Int64
			}
		case file.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case file.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
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
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("folder_id=")
	builder.WriteString(f.FolderID)
	builder.WriteString(", ")
	builder.WriteString("md5=")
	builder.WriteString(f.Md5)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(f.Name)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", f.Size))
	builder.WriteString(", ")
	builder.WriteString("content_type=")
	builder.WriteString(f.ContentType)
	builder.WriteString(", ")
	builder.WriteString("downloads=")
	builder.WriteString(fmt.Sprintf("%v", f.Downloads))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
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

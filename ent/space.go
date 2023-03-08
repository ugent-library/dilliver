// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ugent-library/deliver/ent/space"
)

// Space is the model entity for the Space schema.
type Space struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Admins holds the value of the "admins" field.
	Admins []string `json:"admins,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SpaceQuery when eager-loading is set.
	Edges SpaceEdges `json:"edges"`
}

// SpaceEdges holds the relations/edges for other nodes in the graph.
type SpaceEdges struct {
	// Folders holds the value of the folders edge.
	Folders []*Folder `json:"folders,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FoldersOrErr returns the Folders value or an error if the edge
// was not loaded in eager-loading.
func (e SpaceEdges) FoldersOrErr() ([]*Folder, error) {
	if e.loadedTypes[0] {
		return e.Folders, nil
	}
	return nil, &NotLoadedError{edge: "folders"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Space) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case space.FieldAdmins:
			values[i] = new([]byte)
		case space.FieldID, space.FieldName:
			values[i] = new(sql.NullString)
		case space.FieldCreatedAt, space.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Space", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Space fields.
func (s *Space) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case space.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				s.ID = value.String
			}
		case space.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case space.FieldAdmins:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field admins", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Admins); err != nil {
					return fmt.Errorf("unmarshal field admins: %w", err)
				}
			}
		case space.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case space.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryFolders queries the "folders" edge of the Space entity.
func (s *Space) QueryFolders() *FolderQuery {
	return NewSpaceClient(s.config).QueryFolders(s)
}

// Update returns a builder for updating this Space.
// Note that you need to call Space.Unwrap() before calling this method if this Space
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Space) Update() *SpaceUpdateOne {
	return NewSpaceClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Space entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Space) Unwrap() *Space {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Space is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Space) String() string {
	var builder strings.Builder
	builder.WriteString("Space(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("admins=")
	builder.WriteString(fmt.Sprintf("%v", s.Admins))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Spaces is a parsable slice of Space.
type Spaces []*Space

// Code generated by ent, DO NOT EDIT.

package space

const (
	// Label holds the string label denoting the space type in the database.
	Label = "space"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeFolders holds the string denoting the folders edge name in mutations.
	EdgeFolders = "folders"
	// Table holds the table name of the space in the database.
	Table = "spaces"
	// FoldersTable is the table that holds the folders relation/edge.
	FoldersTable = "folders"
	// FoldersInverseTable is the table name for the Folder entity.
	// It exists in this package in order to avoid circular dependency with the "folder" package.
	FoldersInverseTable = "folders"
	// FoldersColumn is the table column denoting the folders relation/edge.
	FoldersColumn = "space_folders"
)

// Columns holds all SQL columns for space fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)
